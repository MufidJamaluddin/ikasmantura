import React, {PureComponent} from "react";

import {Card, Alert} from "react-bootstrap";
import Image from "../component/Image";
import {ThemeContext} from "../component/PageTemplate";
import {getOrganizations} from "../models/OrganizationModel";

import blankAvatar from './../../resource/blank_avatar.svg'
import ShareSocialMedia from "../component/ShareSocialMedia";

interface OrganizationViewState {
    data: Array<any>|null;
}

function PeopleProfile(props) {
    let {
        id, image, name, userFullname
    } = props

    return (
        <div className="col-auto mb-2" key={id}>
            <Card style={{'width':'12rem'}} className="h-100">
                <Image
                    className="card-img-top"
                    src={image ?? blankAvatar}
                    fallbackSrc={blankAvatar}
                    alt={name}/>
                <Card.Body>
                    <p className={"lead"}>{userFullname ?? 'Anonim'}</p>
                    <b>{name}</b>
                </Card.Body>
            </Card>
        </div>
    )
}

export default class OrganizationView
    extends PureComponent<any,OrganizationViewState>
{
    constructor(props:any) {
        super(props);
        this.state = {
            data: null,
        }
    }

    async updateOrganizationData() {
        let data = await getOrganizations()

        this.setState({ data: data?.data })
    }

    static contextType = ThemeContext;

    componentDidMount()
    {
        this.updateOrganizationData()
        this.context.setHeader({ title: 'Struktur Organisasi', showTitle: true })
    }

    render() {
        let data: Array<any>|null = this.state.data

        if(!Array.isArray(data)) {
            return <div className="c-center-box c-loader"/>
        }

        return (
            <div className="row justify-content-center mb-3">
                {
                    data.length === 0 && (
                        <Alert variant="warning">
                            <strong>Belum ada kepengurusan IKA SMAN Situraja</strong>
                        </Alert>
                    )
                }
                {
                    data.map((item, key) => (
                        <PeopleProfile {...item} key={item.id} />
                    ))
                }
                {
                    data.length > 0 && (
                        <ShareSocialMedia title={'Struktur Organisasi IKA SMAN Situraja'} className="fa-pull-right" />
                    )
                }
            </div>
        )
    }
}
