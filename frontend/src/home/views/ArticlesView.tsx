import * as React from "react";
import {Link} from "react-router-dom";
import {Pagination, Sort} from "ra-core/src/types";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';

import ReactPaginate from 'react-paginate';

interface ArticleItem {
    title?: string
    createdAt_gte?: string|Date,
    createdAt_lte?: string|Date
}

interface ArticlesViewState
{
    total: number;
    pagination: Pagination;
    sort: Sort;
    filter: ArticleItem|any;
    data: Array<any>;
    topics: Array<any>;
    loading: boolean;
}

export default class ArticlesView
    extends React.PureComponent<{}, ArticlesViewState>
{
    constructor(props:any)
    {
        super(props);

        this.state = {
            pagination: {
                page: 1,
                perPage: 3,
            },
            sort: {
                field: 'id',
                order: 'ASC'
            },
            filter: {

            },
            loading: true,
            data: [],
            total: 0,
            topics: [],
        }

        this.handlePageChange = this.handlePageChange.bind(this)
        this.handleSearchForm = this.handleSearchForm.bind(this)
    }

    updateData()
    {
        if(this.state.loading)
        {
            const dataProvider = DataProviderFactory.getDataProvider()

            dataProvider.getList("articles", this.state).then(resp => {
                this.setState(state => {
                    let newState = {
                        data: resp.data,
                        loading: false,
                        total: resp.total,
                    }
                    return {...state, ...newState}
                })
            }, error => {
                NotificationManager.error(error, 'Get Data Error');

                this.setState(state => {
                    let newState = {loading: false}
                    return {...state, ...newState}
                })
            })
        }
    }

    updateTopics()
    {
        let dataProvider = DataProviderFactory.getDataProvider()
        dataProvider.getList("article_topics", {
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
                return {...state, topics: resp.data }
            })
        }, error => {
            NotificationManager.error(error, 'Get Data Error');
        })
    }

    componentDidMount()
    {
        this.updateData()


    }

    componentDidUpdate(prevProps: Readonly<{}>, prevState: Readonly<ArticlesViewState>, snapshot?: any)
    {
        if(
            prevState.filter !== this.state.filter
            || prevState.pagination !== this.state.pagination
        )
        {
            this.updateData()
        }
    }

    handlePageChange(data)
    {
        this.setState(state => {
            let newState = {
                loading: true, pagination: {
                    page: data.selected + 1, perPage: state.pagination.perPage
                }
            }
            return {...state, ...newState}
        })
    }

    handleSearchForm(e)
    {
        try
        {
            e.preventDefault()
        }
        catch (err)
        {

        }

        let formData = new FormData(e.target);
        let title = formData.get('title')

        let filter = {}

        if(title)
        {
            filter = {
                title: title
            }
        }

        this.setState(state => {
            let newState = {
                loading: true, filter: filter, pagination: {
                    page: 1,
                    perPage: state.pagination.perPage
                }
            }

            return {...state, ...newState}
        })

        return false
    }

    render()
    {
        return (
            <>
                <div className={"c-g-banner c-text-center"}>
                    <h1 className={"lead"}>Artikel</h1>
                </div>
                <div className={"c-container c-row"}>
                    <div className={"c-right-column"}>
                        <p className={"lead"}>Kategori</p>
                        <ul className={"c-list-unstyled"}>
                            {
                                this.state.topics.map((item, key) => (
                                    <li key={key}>
                                        <button className={"c-button primary c-margin-bs"}>
                                            {item.name}
                                        </button>
                                    </li>
                                ))
                            }
                        </ul>
                    </div>
                    <div className={"c-left-column"}>
                        <div className={"c-row"}>
                            <form onSubmit={this.handleSearchForm} className={"c-filter"}>
                                <input type="text"
                                       name="title"
                                       placeholder="Judul Berita yang Anda Cari" />

                                <button type="submit" className="c-button primary">
                                    Cari
                                </button>
                            </form>

                            <ReactPaginate
                                previousLabel={'previous'}
                                nextLabel={'next'}
                                breakLabel={'...'}
                                breakClassName={'break-me'}
                                pageCount={ this.state.total / this.state.pagination.perPage }
                                marginPagesDisplayed={2}
                                pageRangeDisplayed={5}
                                onPageChange={this.handlePageChange}
                                containerClassName={'pagination'}
                                subContainerClassName={'pages pagination'}
                                activeClassName={'active'}
                            />
                        </div>
                        {
                            this.state.data.map(item => {
                                return <div key={item.id} className="c-card c-card-large">
                                    <img className={"c-img-full"}
                                         src={item.image ?? "/static/img/jakarta.jpg"}
                                         alt="Avatar"/>
                                    <div className="c-container">
                                        <h4><b>{item.title}</b></h4>
                                        <p>{item.description} &nbsp;
                                            <Link to={`news/${item.id}`}><small><b>Read More...</b></small></Link>
                                        </p>
                                    </div>
                                </div>
                            })
                        }
                    </div>
                </div>
            </>
        )
    }
}