import React from "react";
import RegeImage from "./RegeImage";

import './RegeTitle.css'

export default function RegeTitle(props: any) {
    return (
        <div className="rege-title">
            {props.children}
            <div className="rege-title-color">
                <RegeImage/>
            </div>
        </div>
    )
}
