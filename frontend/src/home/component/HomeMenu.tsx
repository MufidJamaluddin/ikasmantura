import {Button, OverlayTrigger, Tooltip} from "react-bootstrap";
import {Link} from "react-router-dom";
import React from "react";

export default function HomeMenu (props) {

    let {
        id, description, icon, name, linkTo
    } = props

    return (
        <OverlayTrigger
            key={id}
            placement={'bottom'}
            overlay={
                <Tooltip id={`tooltip-${id}`}>
                    {description}
                </Tooltip>
            }>
            {({ ref, ...triggerHandler }) => (
                <div {...triggerHandler}
                     className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                    <div
                        ref={ref}
                        className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                        <i className={`${icon} m-auto`}/>
                    </div>
                    <h3 className="text-success">{name}</h3>
                    <Link to={linkTo}>
                        <Button variant="outline-primary mb-1">
                            Buka
                        </Button>
                    </Link>
                </div>
            )}
        </OverlayTrigger>
    )
}
