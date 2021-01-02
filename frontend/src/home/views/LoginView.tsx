import React, {useState, useEffect, useRef, useContext} from "react";
import {Button, Card, Col, Form, Row} from "react-bootstrap";
import {Link} from "react-router-dom";

import {FieldFeedback, FieldFeedbacks, FormWithConstraints} from "react-form-with-constraints";
import {NotificationManager} from 'react-notifications';
import AuthProvider from "../../dataprovider/authProvider";
import {ThemeContext} from "../component/PageTemplate";
import {login} from "../models/AccountModel";

function FormLogin(props) {

    const formEl = useRef(null);

    async function handleChange({ target }) {
        // Validates only the given fields and returns Promise<Field[]>
        await formEl?.current?.validateFields(target);
    }

    async function onSubmit(e: React.FormEvent<HTMLFormElement>|any) {
        e.preventDefault();

        await formEl?.current?.validateForm();

        if (!(formEl?.current?.isValid())) {
            NotificationManager.error('Mohon Isi Username dan Password dengan Benar!', 'Login Gagal');
            return
        }

        let formData = new FormData(e.target)

        let result = await login({
            username: formData.get('username'),
            password: formData.get('password')
        })
        if(result) {
            this.props.history.push('/panel')
        }
    }

    return (
        <FormWithConstraints
            ref={formEl}
            onSubmit={onSubmit}
            noValidate>

            <Form.Group as={Row} controlId="formUsername">
                <Form.Label srOnly>Username</Form.Label>
                <Form.Control type="text"
                              placeholder="Username"
                              name="username"
                              autoComplete="off"
                              maxLength={35}
                              minLength={3}
                              required={true}
                              onChange={handleChange}
                />
                <FieldFeedbacks for="username">
                    <FieldFeedback when="valueMissing" error className="text-error">
                        Wajib diisi!
                    </FieldFeedback>
                    <FieldFeedback when="tooShort" error className="text-error">
                        Username yang diisikan terlalu pendek!
                    </FieldFeedback>
                    <FieldFeedback when="*" className="text-error" />
                </FieldFeedbacks>
            </Form.Group>

            <Form.Group as={Row} controlId="formPassword">
                <Form.Label srOnly>Password</Form.Label>
                <Form.Control type="password"
                              placeholder="Password"
                              name="password"
                              maxLength={35}
                              required={true}
                              onChange={handleChange}
                />
                <FieldFeedbacks for="password">
                    <FieldFeedback when="valueMissing" error className="text-error">
                        Wajib diisi!
                    </FieldFeedback>
                    <FieldFeedback when="*" className="text-error" />
                </FieldFeedbacks>
            </Form.Group>

            <div className="row justify-content-center">
                <Link to={"/"}>
                    <Button variant="warning" type="button">
                        Kembali
                    </Button>
                </Link>

                &nbsp;
                <Button variant="primary" type="submit">
                    Masuk
                </Button>

            </div>

        </FormWithConstraints>
    )
}

export default function LoginView({ history, ...props })
{
    const [loading, setLoading] = useState(true);
    const theme = useContext(ThemeContext);

    useEffect( () => {
        AuthProvider.checkAuth()
            .then(() => {
                history.replace('/panel')
            })
            .catch(() => {
                setLoading(false)
            })

        theme.setHeader({ title: 'Login', showTitle: false })
    }, []);

    async function onSubmit({ username, password }) {
        history.replace({
            pathname: '/panel/login',
            state: {
                username: username,
                password: password
            }
        })
    }

    return (
        <Row className="d-table h-100 w-100">
            <Col className="d-table-cell align-middle w-100">
                <Card className="col-md-4 col-sm-8 mx-auto">
                    {
                        loading ? (
                            <Card.Text>Loading...</Card.Text>
                        ) : [
                            <Card.Title>
                                <h1 className="text-center">Login</h1>
                            </Card.Title>,
                            <Card.Body>
                                <FormLogin onSubmit={onSubmit} />
                            </Card.Body>,
                            <Card.Footer className="row justify-content-center">
                                <p>
                                Belum mempunyai akun?
                                </p>
                                &nbsp;
                                <Link to={"/register"}>
                                    <Button type="button" variant="warning" size="sm">
                                        Daftar Akun
                                    </Button>
                                </Link>
                            </Card.Footer>,
                        ]
                    }
                </Card>
            </Col>
        </Row>
    )
}
