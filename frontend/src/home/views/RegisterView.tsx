import React, {PureComponent} from "react";
import {Button, Card, Col, Form, Row, Tabs} from "react-bootstrap";
import {Link, RouteComponentProps} from "react-router-dom";

import {NotificationManager} from 'react-notifications';
import {Async, FieldFeedback, FieldFeedbacks, FormWithConstraints} from "react-form-with-constraints";

import Select from "react-select";
import makeAnimated from 'react-select/animated';
import {TIForm, TIFormType} from "../component/CForm";
import {ValidateEmail} from "../../utils/Form";
import {ThemeContext} from "../component/PageTemplate";
import {checkAvailability, registerNewAccount} from "../models/AccountModel";
import {getClassrooms} from "../models/ClassroomsModel";

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function isDeviceSmall(): boolean {
    let width = (window.innerWidth > 0) ? window.innerWidth : window.screen?.width
    return width < 680
}

const selectAnimatedComponents = makeAnimated();

const EmailPattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export default class RegisterView extends PureComponent<RouteComponentProps<any>|any,
    {classrooms: Array<any>, formType: TIFormType, formData: any, inputValidateQueue: any, nextValidation: number}>
{
    private readonly formElement: React.RefObject<FormWithConstraints>
    private inputPassword: HTMLInputElement | null

    static contextType = ThemeContext;

    constructor(props:any)
    {
        super(props);

        let isSmall = isDeviceSmall()

        this.state = {
            classrooms: [],
            formType: isSmall ? TIFormType.INLINE : TIFormType.TABBED,
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

        let formData = this.state.formData

        formData.classrooms = formData.classrooms.map(item => {
            return item.value
        }).map(item => {
            if(item) return parseInt(item)
            return null
        }).filter(item => {
            return item !== null && !isNaN(item)
        })

        let result = await registerNewAccount(formData)
        if(result) {
            this.props.history.push('/login')
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

    async handleChange(e, name)
    {
        let form = this.formElement.current
        if(!form) return

        try
        {
            if(Array.isArray(e)) {
                this.setState(state => ({
                    ...state,
                    formData: {
                        ...state.formData,
                        [name]: e,
                    }
                }))
                return
            }

            if(e.cancelable) {
                e.preventDefault()
            }

            let inputValidateQueue = this.state.inputValidateQueue
            inputValidateQueue[name] = e.target

            let value = e.target.value

            let names = name.split('.')

            let data = {}

            if(names.length === 2) {
                data = {...this.state.formData}
                if(data[names[0]] === null || data[names[0]] === undefined) {
                    data[names[0]] = {}
                }
                data[names[0]][names[1]] = value
            } else {
                data[names[0]] = value
            }

            this.setState(state => ({
                ...state,
                formData: {
                    ...state.formData,
                    ...data
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

    async updateClassrooms()
    {
        let classrooms = await getClassrooms() ?? []

        let nClassrooms = classrooms.map(item => {
            let label
            if(item.major) {
                label = `${item.level}-${item.major}-${item.seq}`
            } else {
                label = `${item.level}-${item.seq}`
            }

            return {
                value: item.id,
                label: label,
            }
        })

        this.setState(state => {
            return {...state, classrooms: nClassrooms }
        })
    }

    componentDidMount()
    {
        this.updateClassrooms()

        this.context.setHeader({ title: 'Pendaftaran Alumni', showTitle: false })
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

                <Row>

                <Col md={6}>
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
                                      onChange={(e) => this.handleChange(e, 'username')}
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
                </Col>

                <Col md={6}>
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
                                      onChange={(e) => this.handleChange(e, 'email')}
                        />
                        <FieldFeedbacks for="email">
                            <FieldFeedback when="valueMissing" error className="text-error">
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when={value => !EmailPattern.test(value)} error className="text-error">
                                Format email tidak valid!
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formPassword">
                    <Form.Label column sm={4}>Password</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="password"
                                      placeholder="Password"
                                      name="password"
                                      value={formData.password}
                                      maxLength={35}
                                      minLength={5}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'password')}
                                      ref={(ref) => this.inputPassword = ref}
                        />
                        <FieldFeedbacks for="password">
                            <FieldFeedback when="valueMissing" error className="text-error">
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="tooShort" error className="text-error">
                                Password yang anda pilih terlalu pendek!
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error className="text-error">
                                Minimal lima karakter
                            </FieldFeedback>
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>
                </Col>

                <Col md={6}>
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
                                      onChange={(e) => this.handleChange(e,'confirmPassword')}
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
                </Col>

                </Row>

            </TIForm>
        )
    }

    renderPersonal({ eventKey, title, type })
    {
        let formData = this.state.formData

        return (
            <TIForm eventKey={eventKey} title={title} type={type}>

                <br/>

                <Row>

                <Col md={6}>
                <Form.Group as={Row} controlId="formName">
                    <Form.Label column sm={4}>Nama Lengkap</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Nama Lengkap"
                                      name="name"
                                      value={formData.name}
                                      minLength={3}
                                      maxLength={35}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'name')}
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formHP">
                    <Form.Label column sm={4}>Nomor HP</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Nomor HP"
                                      name="phone"
                                      value={formData.phone}
                                      minLength={10}
                                      maxLength={13}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'phone')}
                        />
                        <FieldFeedbacks for="phone">
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formJob">
                    <Form.Label column sm={4}>Pekerjaan</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Pekerjaan"
                                      name="job"
                                      value={formData.job}
                                      minLength={2}
                                      maxLength={35}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'job')}
                        />
                        <FieldFeedbacks for="job">
                            <FieldFeedback when="valueMissing" error>
                                Wajib diisi
                            </FieldFeedback>
                            <FieldFeedback when="patternMismatch" error>
                                Minimal isi tiga karakter dan maksimal 85 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formJobDesc">
                    <Form.Label column sm={4}>Deskripsi Pekerjaan</Form.Label>
                    <Col sm={8}>
                        <Form.Control as="textarea" rows={2}
                                      placeholder="Deskripsi Pekerjaan"
                                      name="jobDesc"
                                      value={formData.jobDesc}
                                      minLength={3}
                                      maxLength={85}
                                      required={false}
                                      onChange={(e) => this.handleChange(e,'jobDesc')}
                                      style={{resize: 'none'}}
                        />
                        <FieldFeedbacks for="jobDesc">
                            <FieldFeedback when="patternMismatch" error>
                                Minimal isi tiga karakter dan maksimal 85 karakter
                            </FieldFeedback>
                            <FieldFeedback when="*" className="text-error" />
                        </FieldFeedbacks>
                    </Col>
                </Form.Group>
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formForceYear">
                    <Form.Label column sm={4}>Tahun Lulus</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Tahun Lulus"
                                      name="forceYear"
                                      value={formData.forceYear}
                                      minLength={4}
                                      maxLength={4}
                                      required={true}
                                      pattern="^[0-9]{4}$"
                                      onChange={(e) => this.handleChange(e,'forceYear')}
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formKelas">
                    <Form.Label column sm={4}>Kelas SMA yang Pernah Dijalani</Form.Label>
                    <Col sm={8}>
                        <Form.Control
                            as={Select}
                            closeMenuOnSelect={false}
                            components={selectAnimatedComponents}
                            className="no-padding no-border small"
                            options={this.state.classrooms}
                            isClearable={true}
                            isMulti={true}
                            isLoading={
                                this.state.classrooms.length === 0
                            }
                            placeholder="Kelas"
                            value={formData.classrooms ?? []}
                            onChange={(e) => this.handleChange(e,'classrooms')}
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
                </Col>

                </Row>

            </TIForm>
        )
    }

    renderAddress({ eventKey, title, type, isSmall })
    {
        let formData = this.state.formData
        return (
            <TIForm eventKey={eventKey} title={title} type={type}>

                <br/>

                <Row>

                <Col md={12}>
                <Form.Group as={Row} controlId="formStreet">
                    <Form.Label column sm={2}>Jalan</Form.Label>
                    <Col sm={10}>
                        <Form.Control as="textarea" rows={isSmall ? 2 : 1}
                                      placeholder="Jalan"
                                      name="address.street"
                                      value={formData.address?.street}
                                      minLength={5}
                                      maxLength={75}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'address.street')}
                                      style={{resize: 'none'}}
                        />
                        <FieldFeedbacks for="address.street">
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formSuite">
                    <Form.Label column sm={4}>Kecamatan</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kecamatan"
                                      name="address.suite"
                                      value={formData.address?.suite}
                                      minLength={5}
                                      maxLength={53}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'address.suite')}
                        />
                        <FieldFeedbacks for="address.suite">
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formCity">
                    <Form.Label column sm={4}>Kabupaten/Kota</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kota tempat tinggal anda"
                                      name="address.city"
                                      value={formData.address?.city}
                                      minLength={3}
                                      maxLength={35}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'address.city')}
                        />
                        <FieldFeedbacks for="address.city">
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formZipCode">
                    <Form.Label column sm={4}>Kode Pos</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Kode Pos"
                                      name="address.zipcode"
                                      value={formData.address?.zipcode}
                                      minLength={3}
                                      maxLength={11}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'address.zipcode')}
                        />
                        <FieldFeedbacks for="address.zipcode">
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
                </Col>

                <Col md={6}>
                <Form.Group as={Row} controlId="formNation">
                    <Form.Label column sm={4}>Negara</Form.Label>
                    <Col sm={8}>
                        <Form.Control type="text"
                                      placeholder="Negara"
                                      name="address.state"
                                      value={formData.address?.state}
                                      minLength={5}
                                      maxLength={16}
                                      required={true}
                                      onChange={(e) => this.handleChange(e,'address.state')}
                        />
                        <FieldFeedbacks for="address.state">
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
                </Col>

                </Row>

            </TIForm>
        )
    }

    render()
    {
        let formType = this.state.formType

        let isSmall = isDeviceSmall()

        let FormParent

        if(isSmall) {
            FormParent = Row
        } else {
            FormParent = formType === TIFormType.TABBED ? Tabs : Row
        }

        return (
            <Row>
                <Col lg={{span:10, offset:1}}>
                    <Card className={formType === TIFormType.TABBED ? 'abu' : ''}>

                        {
                            (!isSmall) && (
                                <Card.Body>
                                    <div className="fa-pull-left">
                                        Pilih tampilan &nbsp;
                                        <Button
                                            className={
                                                formType === TIFormType.TABBED ?
                                                    'btn-kegiatan-active' : 'btn-kegiatan'
                                            }
                                            size="sm"
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
                                            size="sm"
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
                            )
                        }

                        <Card.Title className="text-center tabTitle">
                            Pendaftaran Alumni SMAN Situraja
                        </Card.Title>

                        <FormWithConstraints
                            ref={this.formElement}
                            onSubmit={this.onSubmit}
                            noValidate>

                            <Card.Body>

                                <FormParent className="myTab"
                                            defaultActiveKey="account"
                                            id="uncontrolled-forms">

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
                                            type: formType,
                                            isSmall: isSmall
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
        )
    }
}
