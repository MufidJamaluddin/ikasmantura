import React, {PureComponent} from "react";

import ImageGallery from 'react-image-gallery';

import 'react-image-gallery/styles/css/image-gallery.css'
import {Alert, Badge, Button, Col, Row} from "react-bootstrap";

import {NotificationManager} from 'react-notifications';
import {ThemeContext} from "../component/PageTemplate";
import {Album, getAlbums} from "../models/AlbumModel";
import {getPhotoByAlbumIds, Photo} from "../models/PhotoModel";

interface GalleryViewState {
    albums: Array<Album>
    selected: Array<Album>
    photos: Array<Photo>
    isLoading: boolean
}

export default class GalleryView extends PureComponent<any, GalleryViewState>
{
    static contextType = ThemeContext;

    constructor(props: any)
    {
        super(props);
        this.state = {
            albums: [],
            selected: [],
            photos: [],
            isLoading: true
        }

        this.chooseAlbum = this.chooseAlbum.bind(this)
    }

    async updateAlbums() {
        let albums = await getAlbums() ?? []

        this.setState(oldState => {
            return {...oldState, albums: albums as Array<Album>, isLoading: false}
        })
    }

    async updatePhotos() {
        if (this.state.isLoading) {
            let requestedAlbumIds = this.state.selected.map(item => item.id)

            let photos = await getPhotoByAlbumIds(requestedAlbumIds) ?? []

            this.setState(state => {
                let newState = {
                    photos: photos as Array<Photo>,
                    isLoading: false,
                }
                return {...state, ...newState}
            })
        }
    }

    componentDidMount()
    {
        try
        {
            this.updateAlbums();
            this.context.setHeader({ title: 'Galeri', showTitle: true })
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    componentDidUpdate(prevProps: Readonly<{}>, prevState: Readonly<{}>, snapshot?: any)
    {
        try
        {
            this.updatePhotos()
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    chooseAlbum(e: any, id: number|string)
    {
        e.preventDefault()

        let selected = this.state.albums.filter(item => item.id === id)
        selected = selected.concat(this.state.selected)

        let selectedIds = selected.map(item => item.id)

        let currentAlbums = this.state.albums.map(item => {
            if(selectedIds.includes(item.id)) return null
            else return item
        }).filter(item => item !== null)

        this.setState(oldState => {
            return {...oldState, isLoading: true, selected: selected, albums: currentAlbums}
        })
    }

    unChooseAlbum(e: any, id: number|string)
    {
        e.preventDefault()

        let albums = this.state.selected.filter(item => item.id === id)
        albums = albums.concat(this.state.albums)

        let albumIds = albums.map(item => item.id)

        let currentSelected = this.state.selected.map(item => {
            if(albumIds.includes(item.id)) return null
            else return item
        }).filter(item => item !== null)

        this.setState(oldState => {
            return {...oldState, isLoading: true, selected: currentSelected, albums: albums, photos: []}
        })
    }

    renderAlbums()
    {
        return (
            <Col md={2}>
                <h5>Pilih Album</h5>
                <p>Mohon klik album yang akan anda pilih</p>
                <ul className="list-unstyled">
                    {
                        this.state.albums.length > 0 ?
                            this.state.albums.map(item => <li key={item.id} className="mb-1">
                                <Button variant="primary" onClick={e => this.chooseAlbum(e, item.id)}>
                                    {item.title}
                                </Button>
                            </li>) : (
                                <Alert variant="warning">
                                    Semua album telah dipilih!
                                </Alert>
                            )
                    }
                </ul>
            </Col>
        )
    }

    renderSelectedAlbums()
    {
        return (
            <Col md={2}>
                <h5>Album Saat Ini</h5>
                <p>Album yang sedang ditampilkan</p>
                <ul className="list-unstyled">
                    {
                        this.state.selected.length > 0 ?
                            this.state.selected.map(item => <li key={item.id} className="mb-1">
                                <div className="btn btn-info"
                                     onClick={e => this.unChooseAlbum(e, item.id)}>
                                    {item.title}
                                    &nbsp;
                                    <Badge variant="danger">
                                        <span className="fas fa-times fa-1x"/>
                                    </Badge>
                                </div>
                            </li>) : (
                                <Alert variant="warning">
                                    Anda belum memilih album!
                                </Alert>
                            )
                    }
                </ul>
            </Col>
        )
    }

    render()
    {
        if(this.state.isLoading) return (
            <Row className="justify-content-center mb-3">
                Loading Bro...
            </Row>
        )

        let isWidthWide = (window.innerWidth <= 760);

        return (
            <Row>
                {
                    this.renderAlbums()
                }
                {
                    isWidthWide && this.renderSelectedAlbums()
                }
                <Col md={8}>
                    {
                        this.state.photos.length > 0 ? (
                            <ImageGallery items={this.state.photos}/>
                        ) : (
                            <Alert variant="warning">
                                Album yang dipilih tidak memiliki foto, coba pilih album lain!
                            </Alert>
                        )
                    }
                </Col>
                {
                    !isWidthWide && this.renderSelectedAlbums()
                }
            </Row>
        )
    }
}
