import React, {PureComponent} from "react";
import RegeTitle from "../component/RegeTitle";
import {Button, Card, Col, Container, Form, Row, Tab, Tabs} from "react-bootstrap";
import {Link, RouteComponentProps} from "react-router-dom";

import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import {Async, FieldFeedback, FieldFeedbacks, FormWithConstraints} from "react-form-with-constraints";

import Select from "react-select";
import makeAnimated from 'react-select/animated';

const selectAnimatedComponents = makeAnimated();

/*
function sleep(ms: number)
{
    return new Promise(resolve => setTimeout(resolve, ms));
}
 */

async function checkAvailability({ username = '', email = '' })
{
    /*
    if(
        process.env.NODE_ENV === "development"
        || process.env.NODE_ENV === "test"
    )
    {
        await sleep(1000);
        return true;
    }
     */

    let dataProvider = DataProviderFactory.getDataProvider()
    let result:boolean

    try
    {
        result = await dataProvider.create('register/availability', {
            data: {
                username: username,
                email: email,
            }
        }).then((resp: any) => {
            return resp.exist
        }, error => {
            NotificationManager.error(error.message, error.name);
            return false
        });
    }
    catch (e)
    {
        NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        result = false
    }

    return result
}

export default class RegisterView extends PureComponent<RouteComponentProps<any>, {classrooms: Array<any>}>
{
    private readonly formElement: React.RefObject<FormWithConstraints>

    constructor(props:any)
    {
        super(props);

        this.state = {
            classrooms: [],
        }

        this.formElement = React.createRef()

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

        let serialize = require('form-serialize');

        let data = serialize(e.currentTarget, { hash: true });

        let dataProvider = DataProviderFactory.getDataProvider()

        data.classrooms = data.classrooms.map(item => {
            return { id: item }
        })

        try
        {
            dataProvider.create("temp_users", data).then(_ => {

                NotificationManager.success(
                    'Pendaftaran Sukses, Mohon Tunggu Konfirmasi Admin!', 'Pendaftaran Sukses');

                this.props.history.push('/login')

            }, error => {
                NotificationManager.error(error.message, 'Pendaftaran Gagal');
            })
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    async handleChange(e)
    {
        let form = this.formElement.current
        if(!form) return;

        try
        {
            e.preventDefault()

            await form.validateFields(e.target);
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

    renderAccount()
    {
        return (
            <Tab eventKey="account" title="Data Akun">

                <br/>

                <Form.Group controlId="formUsername">
                    <Form.Label>Username</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Username"
                                  name="username"
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
                </Form.Group>

                <Form.Group controlId="formEmail">
                    <Form.Label>Email</Form.Label>
                    <Form.Control type="email"
                                  placeholder="Email"
                                  name="email"
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
                </Form.Group>

                <Form.Group controlId="formPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password"
                                  placeholder="Password"
                                  name="password"
                                  minLength={5}
                                  maxLength={35}
                                  required={true}
                                  onChange={this.handleChange}
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
                </Form.Group>
            </Tab>
        )
    }

    renderPersonal()
    {
        return (
            <Tab eventKey="profile" title="Data Personal">

                <br/>

                <Form.Group controlId="formName">
                    <Form.Label>Nama</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Nama"
                                  name="name"
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
                </Form.Group>

                <Form.Group controlId="formHP">
                    <Form.Label>Nomor HP</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Nomor HP"
                                  name="hp"
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
                </Form.Group>

                <Form.Group controlId="formAngkatan">
                    <Form.Label>Angkatan</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Angkatan"
                                  name="forceYear"
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
                </Form.Group>

                <Form.Group controlId="formKelas">
                    <Form.Label>Kelas SMAN Situraja yang Pernah Dijalani</Form.Label>
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
                        onChange={this.handleChange}
                        required={true}
                    />
                    <FieldFeedbacks for="classroom">
                        <FieldFeedback when="valueMissing" error>
                            Wajib diisi
                        </FieldFeedback>
                        <FieldFeedback when="*" className="text-error" />
                    </FieldFeedbacks>
                </Form.Group>

            </Tab>
        )
    }

    renderAddress()
    {
        return (
            <Tab eventKey="address" title="Data Alamat">

                <br/>

                <Form.Group controlId="formStreet">
                    <Form.Label>Jalan</Form.Label>
                    <Form.Control as="textarea" rows={2}
                                  placeholder="Jalan"
                                  name="address[street]"
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
                </Form.Group>

                <Form.Group controlId="formSuite">
                    <Form.Label>Kecamatan</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Kecamatan"
                                  name="address[suite]"
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
                </Form.Group>

                <Form.Group controlId="formCity">
                    <Form.Label>Kabupaten / Kota</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Kota tempat tinggal anda"
                                  name="address[city]"
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
                </Form.Group>

                <Form.Group controlId="formZipCode">
                    <Form.Label>Kode Pos</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Kode Pos"
                                  name="address[zipcode]"
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
                </Form.Group>

                <Form.Group controlId="formNation">
                    <Form.Label>Negara</Form.Label>
                    <Form.Control type="text"
                                  placeholder="Negara"
                                  name="address[state]"
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
                </Form.Group>

            </Tab>
        )
    }

    render()
    {
        return (
            <section className="features-icons bg-light">
                <RegeTitle/>
                <Container>
                    <Row>
                        <Col md={{span:10, offset:1}}>
                            <Card>

                                <Card.Body>
                                    Sudah pernah daftar? &nbsp;
                                    <Link to={"/login"}>
                                        <Button variant="warning" size="sm" type="button">
                                            Masuk
                                        </Button>
                                    </Link>
                                </Card.Body>

                                <Card.Title>
                                    <h1 className="text-center">Pendaftaran Alumni</h1>
                                </Card.Title>

                                <FormWithConstraints
                                    ref={this.formElement}
                                    onSubmit={this.onSubmit}
                                    noValidate>

                                    <Card.Body>

                                        <Tabs defaultActiveKey="account" id="uncontrolled-tab-example">
                                            {this.renderAccount()}

                                            {this.renderPersonal()}

                                            {this.renderAddress()}
                                        </Tabs>

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
