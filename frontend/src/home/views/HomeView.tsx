import React from "react";
import {RouteComponentProps} from "react-router";
import {Link} from "react-router-dom";

export default class HomeView extends React.PureComponent<RouteComponentProps>
{
    render() {
        return (
            <div className="c-g-banner c-text-center c-p-full-height">
                <div>
                    <img src={process.env.PUBLIC_URL + "/static/img/logo_ika2.png"}
                        height={"200px"} alt={"Logo IKA Smantura"}
                    />
                </div>
                <div>
                    <h1>Ikatan Alumni <br/> SMA Negeri Situraja</h1>
                    <p className={"lead"}>
                        <i>Bersatu Kita Teguh...</i>
                    </p>
                </div>
                <div>
                    <Link to="/about">
                        <div className="c-button primary c-margin-bs">
                            Apa itu Ikatan Alumni SMAN Situraja?
                        </div>
                    </Link>
                    <Link to="/organization">
                        <div className="c-button primary c-margin-bs">
                            Struktur Organisasi
                        </div>
                    </Link>
                </div>
                <div>
                    <Link to="/news">
                        <div className="c-button primary c-margin-bs">
                            Berita Terbaru
                        </div>
                    </Link>
                    <Link to="/events">
                        <div className="c-button primary c-margin-bs">
                            Agenda Terbaru
                        </div>
                    </Link>
                    <Link to="/panel">
                        <div className="c-button primary c-margin-bs">
                            Masuk
                        </div>
                    </Link>
                </div>
            </div>
        )
    }
}