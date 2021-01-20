import React, {useContext, useEffect} from "react";
import { Col, Row } from "react-bootstrap";
import {useStore} from "../models";
import HomeMenu from "../component/HomeMenu";
import {ThemeContext} from "../component/PageTemplate";
import {strip_tags} from "../../utils/Security";

export default function HomeView(props)
{
    function getHeader() {
        return [
            <p className="h1 text-center text-white-50 font-weight-bold">
                Selamat Datang di
            </p>,
            <div className="divider-custom">
                <div className="divider-custom-line"/>
                <div className="divider-custom-icon">
                    <i className="fa fa-medal"/>
                </div>
                <div className="divider-custom-line"/>
            </div>,
            <p className="display-4 text-center text-white-50 font-weight-bold">Ikatan Alumni SMAN Situraja</p>,
        ]
    }

    const [state, actions] = useStore('ArticleTopicModel')
    const theme = useContext(ThemeContext);

    useEffect(() => {
        theme.setHeader({ header: getHeader(), title: 'IKA SMAN Situraja', showTitle: false })
        // @ts-ignore
        actions.init()
    }, [])

    let topics = state.data ?? []

    return (
        <Row className="text-center justify-content-center">

            <Col sm={6} md={4} lg={2}>
                <HomeMenu
                    id={'events'}
                    name={'Kegiatan'}
                    icon={'fas fa-calendar-alt text-danger'}
                    description={'Kegiatan yang akan dan telah dilaksanakan'}
                    linkTo={'/events'}
                />
            </Col>

            <Col sm={6} md={4} lg={2}>
                <HomeMenu
                    id={'gallery'}
                    name={'Galleri'}
                    icon={'fas fa-images text-secondary'}
                    description={'Galeri foto acara dan kegiatan'}
                    linkTo={'/gallery'}
                />
            </Col>

            {
                topics.map(item => {
                    let description = strip_tags(item.description)
                    return (
                        <Col sm={6} md={4} lg={2} key={item.id}>
                            <HomeMenu
                                id={item.id}
                                name={item.name}
                                icon={item.icon}
                                description={description}
                                linkTo={{
                                    pathname: '/articles',
                                    state: {
                                        topicId: item.id
                                    }
                                }}
                            />
                        </Col>
                    )
                })
            }

        </Row>
    )
}
