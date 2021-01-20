import React, {PureComponent} from "react";
import {RouteComponentProps} from "react-router-dom";
import {Button, Card, Col, Row} from "react-bootstrap";

import {NotificationManager} from 'react-notifications';
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

async function confirmEmailAccount(username: string, token: string) {
    let dataProvider = DataProviderFactory.getDataProvider()
    try {
        return await dataProvider.create(
            `confirms/tu_emails/${username}/${token}`, { data: {} }).then(_ => {
            NotificationManager.success(
                `Konfirmasi Email Akun ${username} Sukses!`, 'Konfirmasi Sukses')
            return true
        }, error => {
            if(error.status === 403) {
                NotificationManager.error(`Konfirmasi Tidak Valid, Anda Mungkin Telah Konfirmasi Sebelumnya! ${error.message}`, 'Konfirmasi Gagal')
            } else {
                NotificationManager.error(error.message, `Konfirmasi Gagal: ${error.name}`)
            }
            return false
        })
    } catch (e) {
        NotificationManager.error(e.toString(), 'Kesalahan Teknis!');
        return false
    }
}

export class ConfirmEmailView extends PureComponent<
    RouteComponentProps<{username: string, token: string}>>
{
    render()
    {
        let { username = '', token = '' } = this.props.match?.params ?? {}

        let isNotValid = username === '' || token === ''

        const onConfirm = (e) => {
            if(e.cancelable) {
                e.preventDefault()
            }

            confirmEmailAccount(username, token).then(_ => {
                this.props.history.push('/')
            })
        }

        return (
            <Row>
                <Col lg={{span:10, offset:1}}>
                    <Card className="abu">
                        {
                            isNotValid ? (
                                <Card.Title className="text-center tabTitle">
                                    Tidak Valid, Mohon Kembali!
                                </Card.Title>
                            ) : [
                                <Card.Title className="text-center tabTitle">
                                    Konfirmasi Email
                                </Card.Title>,
                                <Card.Body>
                                    Hai {username}, bila akun anda belum dikonfirmasi oleh admin,
                                    mohon klik tombol dibawah ini untuk konfirmasi kepemilikan email.
                                </Card.Body>,
                                <Card.Footer>
                                    <div className="row justify-content-center">
                                        <Button variant="primary" type="button" onClick={onConfirm}>
                                            Kirim Pendaftaran
                                        </Button>
                                    </div>
                                </Card.Footer>
                            ]
                        }
                    </Card>
                </Col>
            </Row>
        )
    }
}
