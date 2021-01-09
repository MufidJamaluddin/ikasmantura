import React from "react";
import {
    EmailIcon,
    EmailShareButton,
    FacebookIcon,
    FacebookShareButton, LineIcon, LineShareButton,
    TwitterIcon,
    TwitterShareButton,
    WhatsappIcon,
    WhatsappShareButton
} from "react-share";

import './ShareSocialMedia.css'

export default function ShareSocialMedia (props) {
    let {
        shareUrl = props.location?.pathname ?? window.location?.href,
        title = 'IKA SMAN Situraja'
    } = props

    return (
        <div className={`Demo__container ${props.className}`}>
            <div className="Demo__some-network">
                <FacebookShareButton
                    url={shareUrl}
                    quote={title}
                    className="Demo__some-network__share-button"
                >
                    <FacebookIcon size={32} round />
                </FacebookShareButton>
            </div>

            <div className="Demo__some-network">
                <TwitterShareButton
                    url={shareUrl}
                    title={title}
                    className="Demo__some-network__share-button"
                    >
                    <TwitterIcon size={32} round />
                </TwitterShareButton>
            </div>

            <div className="Demo__some-network">
                <LineShareButton
                    url={shareUrl}
                    title={title}
                    className="Demo__some-network__share-button"
                >
                    <LineIcon size={32} round />
                </LineShareButton>
            </div>

            <div className="Demo__some-network">
                <WhatsappShareButton
                    url={shareUrl}
                    title={title}
                    separator=":: "
                    className="Demo__some-network__share-button"
                >
                    <WhatsappIcon size={32} round />
                </WhatsappShareButton>
            </div>

            <div className="Demo__some-network">
                <EmailShareButton
                    url={shareUrl}
                    subject={title}
                    body="body"
                    className="Demo__some-network__share-button"
                >
                    <EmailIcon size={32} round />
                </EmailShareButton>
            </div>

        </div>
    );
}
