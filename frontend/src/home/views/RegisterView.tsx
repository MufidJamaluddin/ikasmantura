import React, {PureComponent} from "react";
import RegeTitle from "../component/RegeTitle";
import {Button, Card, Col, Container, Form, Row, Tabs} from "react-bootstrap";
import {Link, RouteComponentProps} from "react-router-dom";

import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import {Async, FieldFeedback, FieldFeedbacks, FormWithConstraints} from "react-form-with-constraints";

import Select from "react-select";
import makeAnimated from 'react-select/animated';
import {TIForm, TIFormType} from "../component/CForm";
import {ValidateEmail} from "../../utils/Form";

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

const selectAnimatedComponents = makeAnimated();

async function checkAvailability({ username = '', email = '' })
{
    let dataProvider = DataProviderFactory.getDataProvider()
    var result:boolean
    var lastUsername
    var lastEmail

    try
    {
        if(lastUsername === username && lastEmail === email)
        {
            return result
        }

        lastUsername = username
        lastEmail = email

        result = await fetch('api/v1/register/availability', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                email: email,
            })
        })
        .then(resp => {
            if(resp.status < 200 && resp.status > 399)
                throw new Error('Server sibuk!')

            return resp.json()
        })
        .then((resp:any) => {
            console.log(resp)
            return !resp?.exist
        }).catch(error => {
            NotificationManager.error(error?.toString(), 'Cek Ketersediaan Akun: Error Koneksi');
            return true
        })

        if(result)
        {
            NotificationManager.info(
                `Akun "${username} ${email}" tersedia!`,
                'Akun tersedia');
        }
        else
        {
            NotificationManager.warning(
            `Akun "${username} ${email}" tidak tersedia!`,
            'Akun tidak tersedia');
        }
    }
    catch (e)
    {
        result = true
    }

    return result
}

export default class RegisterView extends PureComponent<RouteComponentProps<any>,
    {classrooms: Array<any>, formType: TIFormType, formData: any, inputValidateQueue: any, nextValidation: number}>
{
    private readonly formElement: React.RefObject<FormWithConstraints>
    private inputPassword: HTMLInputElement | null

    constructor(props:any)
    {
        super(props);

        this.state = {
            classrooms: [],
            formType: TIFormType.TABBED,
            formData: {},
            inputValidateQueue: {},
            nextValidation: 0
        }

        this.formElement = React.createRef()
        this.inputPassword = null

        this.handleChange = this.handleChange.bind(this)
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e: React.FormEvent<HTMLFormElement>|any)
    {
        e.preventDefault();

        await this.formElement.current.validateForm();

        if(!this.formElement.current.isValid())
        {
            NotificationManager.error('Isian tidak valid! mohon segera perbaiki!', 'Pendaftaran Error');
            return;
        }

        let dataProvider = DataProviderFactory.getDataProvider()
        let formData = this.state.formData

        formData.classrooms = formData.classrooms.map(item => {
            return { id: item }
        })

        try
        {
            dataProvider.create("temp_users", formData).then(_ => {

                NotificationManager.success(
                    'Pendaftaran Sukses, Mohon Tunggu Konfirmasi Admin!', 'Pendaftaran Sukses');

                this.props.history.push('/login')

            }, error => {
                NotificationManager.error(error.message, `Pendaftaran Gagal: ${error.name}`);
            })
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    async handleValidateInput()
    {
        await sleep(2000)

        let form = this.formElement.current
        if(!form) return;

        let validationQueue = this.state.inputValidateQueue

        for (const key of Object.keys(validationQueue)) {
            let target = validationQueue[key]
            await form.validateFields(target)
            delete validationQueue[key]
        }

        this.setState(state => ({
            ...state,
            inputValidateQueue: validationQueue,
        }))
    }

    async handleChange(e)
    {
        let form = this.formElement.current
        if(!form) return;

        try
        {
            e.preventDefault()

            let inputValidateQueue = this.state.inputValidateQueue
            inputValidateQueue[e.target.name] = e.target

            this.setState(state => ({
                ...state,
                formData: {
                    ...state.formData,
                    [e.target.name]: e.target.value,
                },
                inputValidateQueue: inputValidateQueue
            }))
        }
        catch (err)
        {
            console.log(err)
        }
    }

    async checkUsernameAvailability(value: string)
    {
        if(value) {
            if(value.length < 3) return false
        } else return false

        return await checkAvailability({ username: value })
    }

    async checkEmailAvailability(value: string)
    {
        if(value) {
            if(value.length < 5) return false
        } else return false

        if(!ValidateEmail(value)) return false

        return await checkAvailability({ email: value })
    }

    updateClassrooms()
    {
        try
        {
            let dataProvider = DataProviderFactory.getDataProvider()
            dataProvider.getList("classrooms", {
                pagination: {
                    page: 1,
                    perPage: 100,
                },
                sort: {
                    field: 'id',
                    order: 'ASC'
                },
                filter: {
                },
            }).then(resp => {
                this.setState(state => {

                    let optionsData = resp.data.map(item => {
                        let label = `${item.level} - ${item.major} - ${item.seq}`
                        return {
                            value: item.id,
                            label: label,
                        }
                    })

                    return {...state, classrooms: optionsData }
                })
            }, error => {
                NotificationManager.error(error.message, error.name);
            })
        }
        catch(e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
            this.setState(state => {
                return {...state, classrooms: [] }
            })
        }
    }

    componentDidMount()
    {
        this.updateClassrooms()
    }

    componentDidUpdate(prevProps: any, prevState: any, snapshot?: any)
    {
        this.handleValidateInput().then(r => null)
    }

    renderAccount({ eventKey, title, type })
    {
        let formData = this.state.formData
        return (
            <TIForm eventKey={eventKey} title={title} type={type}>

                <br/>

                <Form.Group as={Row} controlId="formUsername">
                    <Form.Label column sm={4}>Username</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Username"
                                      name="username"
                                      value={formData.username}
                                      autoComplete="off"
                                      maxLength={35}
                                      minLength={3}
                                      onChange={this.handleChange}
                                      required={true}
                        />
                        <FieldFeedbacks for="username">
                            <FieldFeedback when="tooShort" error className="text-error">
                                Username yang anda pilih terlalu pendek!
                            </FieldFeedback>
                            <Async
                                promise={this.checkUsernameAvailability}
                                then={available => available ?
                                    <FieldFeedback key="1" info className="text-white">
                                        Username tersedia
                                    </FieldFeedback> :
                                    <FieldFeedback key="2" error className="text-error">
                                        Username telah dimiliki oleh akun lain, mohon pilih username lain!
                                    </FieldFeedback>
                                }
                            />
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formEmail">
                    <Form.Label column sm={4}>Email</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="email"
                                      placeholder="Email"
                                      name="email"
                                      value={formData.email}
                                      autoComplete="off"
                                      minLength={5}
                                      maxLength={250}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="email">
                            <FieldFeedback when={value => !/[^-\s]/.test(value)} error>
                                Tidak boleh mengandung spasi!
                            </FieldFeedback>
                            <Async
                                promise={this.checkEmailAvailability}
                                then={available => available ?
                                    <FieldFeedback key="1" info className="text-white">
                                        Email OK
                                    </FieldFeedback> :
                                    <FieldFeedback key="2" className="text-error">
                                        Email telah dimiliki oleh akun lain, mohon login jika anda mempunyai akun!
                                    </FieldFeedback>
                                }
                            />
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formPassword">
                    <Form.Label column sm={4}>Password</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="password"
                                      placeholder="Password"
                                      name="password"
                                      value={formData.password}
                                      minLength={5}
                                      maxLength={35}
                                      required={true}
                                      onChange={this.handleChange}
                                      ref={(ref) => this.inputPassword = ref}
                        />
                        <FieldFeedbacks for="password">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal lima karakter
                            </FieldFeedback>
                            <FieldFeedback when={value => !/\d/.test(value)} warning>
                                Harus mengandung kombinasi angka
                            </FieldFeedback>
                            <FieldFeedback when={value => !/[a-z]/.test(value)} warning>
                                Harus mengandung kombinasi huruf kecil
                            </FieldFeedback>
                            <FieldFeedback when={value => !/[A-Z]/.test(value)} warning>
                                Harus mengandung kombinasi huruf besar
                            </FieldFeedback>
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formConfirmPassword">
                    <Form.Label column sm={4}>Konfirmasi Password</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="password"
                                      placeholder="Password harus sama"
                                      name="confirmPassword"
                                      value={formData.confirmPassword}
                                      minLength={5}
                                      maxLength={35}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="confirmPassword">
                            <FieldFeedback when={
                                value => value !== this.inputPassword!.value
                            } className="text-error" error>
                                Password tidak sama
                            </FieldFeedback>
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

            </TIForm>
        )
    }

    renderPersonal({ eventKey, title, type })
    {
        let formData = this.state.formData

        return (
            <TIForm eventKey={eventKey} title={title} type={type}>

                <br/>

                <Form.Group as={Row} controlId="formName">
                    <Form.Label column sm={4}>Nama</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Nama"
                                      name="name"
                                      value={formData.name}
                                      minLength={3}
                                      maxLength={35}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="name">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal tiga karakter dan maksimal 35 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formHP">
                    <Form.Label column sm={4}>Nomor HP</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Nomor HP"
                                      name="hp"
                                      value={formData.hp}
                                      minLength={10}
                                      maxLength={13}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="hp">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal 10 karakter dan maksimal 13 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formAngkatan">
                    <Form.Label column sm={4}>Angkatan</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Angkatan"
                                      name="forceYear"
                                      value={formData.forceYear}
                                      minLength={4}
                                      maxLength={4}
                                      required={true}
                                      pattern="^[0-9]{4}$"
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="forceYear">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Wajib menggunakan format tahun dengan benar!
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formKelas">
                    <Form.Label column sm={4}>Kelas SMAN Situraja yang Pernah Dijalani</Form.Label>
                    <Col sm={8}>
                        <Form.Control
                            as={Select}
                            closeMenuOnSelect={false}
                            components={selectAnimatedComponents}
                            className="no-padding no-border"
                            options={this.state.classrooms}
                            isClearable={true}
                            isMulti={true}
                            isLoading={
                                this.state.classrooms.length === 0
                            }
                            placeholder="Kelas"
                            name="classrooms"
                            value={formData.classrooms ?? []}
                            onChange={this.handleChange}
                            required={true}
                        />
                        <FieldFeedbacks for="classroom">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

            </TIForm>
        )
    }

    renderAddress({ eventKey, title, type })
    {
        let formData = this.state.formData
        return (
            <TIForm eventKey={eventKey} title={title} type={type}>

                <br/>

                <Form.Group as={Row} controlId="formStreet">
                    <Form.Label column sm={4}>Jalan</Form.Label>
                    <Col sm={8}>
                        <Form.Control as="textarea" rows={2}
                                      placeholder="Jalan"
                                      name="address[street]"
                                      value={formData.address?.street}
                                      minLength={5}
                                      maxLength={75}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="address[street]">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal isi lima karakter dan maksimal 75 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formSuite">
                    <Form.Label column sm={4}>Kecamatan</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kecamatan"
                                      name="address[suite]"
                                      value={formData.address?.suite}
                                      minLength={5}
                                      maxLength={53}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="address[suite]">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal lima karakter dan maksimal 35 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formCity">
                    <Form.Label column sm={4}>Kabupaten / Kota</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kota tempat tinggal anda"
                                      name="address[city]"
                                      value={formData.address?.city}
                                      minLength={3}
                                      maxLength={35}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="address[city]">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                               Minimal tiga karakter dan maksimal 35 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formZipCode">
                    <Form.Label column sm={4}>Kode Pos</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kode Pos"
                                      name="address[zipcode]"
                                      value={formData.address?.zipcode}
                                      minLength={3}
                                      maxLength={11}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="address[zipcode]">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Wajib menggunakan format kode pos dengan benar!
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formNation">
                    <Form.Label column sm={4}>Negara</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Negara"
                                      name="address[state]"
                                      value={formData.address?.state}
                                      minLength={5}
                                      maxLength={16}
                                      required={true}
                                      onChange={this.handleChange}
                        />
                        <FieldFeedbacks for="address[state]">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Nama negara minimal lima karakter dan maksimal 16 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>

            </TIForm>
        )
    }

    render()
    {
        let formType = this.state.formType

        let FormParent = formType === TIFormType.TABBED ? Tabs : Row;

        return (
            <section className="features-icons bg-light">
                <RegeTitle/>
                <Container>
                    <Row>
                        <Col md={{span:10, offset:1}}>
                            <Card>

                                <Card.Body>
                                    <div className="fa-pull-left">
                                        Pilih tampilan &nbsp;
                                        <Button
                                            className={
                                                formType === TIFormType.TABBED ?
                                                    'btn-kegiatan-active' : 'btn-kegiatan'
                                            }
                                            onClick={(e) => {
                                                e.preventDefault()
                                                this.setState(state => ({
                                                    ...state,
                                                    formType: TIFormType.TABBED
                                                }))
                                            }}
                                        >
                                            Isian Tab Ke Samping
                                        </Button>
                                        <Button
                                            className={
                                                formType === TIFormType.INLINE ?
                                                    'btn-kegiatan-active' : 'btn-kegiatan'
                                            }
                                            onClick={(e) => {
                                                e.preventDefault()
                                                this.setState(state => ({
                                                    ...state,
                                                    formType: TIFormType.INLINE
                                                }))
                                            }}
                                        >
                                            Semua Isian Kebawah
                                        </Button>
                                    </div>
                                    <div className="fa-pull-right">
                                        Sudah pernah daftar? &nbsp;
                                        <Link to={"/login"}>
                                            <Button variant="warning" size="sm" type="button">
                                                Masuk
                                            </Button>
                                        </Link>
                                    </div>
                                </Card.Body>

                                <Card.Title>
                                    <h1 className="text-center">Pendaftaran Alumni</h1>
                                </Card.Title>

                                <FormWithConstraints
                                    ref={this.formElement}
                                    onSubmit={this.onSubmit}
                                    noValidate>

                                    <Card.Body>

                                        <FormParent defaultActiveKey="account" id="uncontrolled-forms">
                                            {this.renderAccount(
                                                {
                                                    title: 'Data Akun',
                                                    eventKey: 'account',
                                                    type: formType
                                                }
                                            )}

                                            {this.renderPersonal(
                                                {
                                                    title: 'Data Pribadi',
                                                    eventKey: 'personal',
                                                    type: formType
                                                }
                                            )}

                                            {this.renderAddress(
                                                {
                                                    title: 'Data Alamat Saat Ini',
                                                    eventKey: 'address',
                                                    type: formType
                                                }
                                            )}
                                        </FormParent>

                                    </Card.Body>
                                    <Card.Footer>
                                        <div className="row justify-content-center">
                                            <Button variant="primary" type="submit">
                                                Kirim Pendaftaran
                                            </Button>
                                        </div>
                                    </Card.Footer>

                                </FormWithConstraints>

                            </Card>
                        </Col>
                    </Row>
                </Container>
            </section>
        )
    }
}
