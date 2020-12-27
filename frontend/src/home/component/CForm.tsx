import {Col, Tab} from "react-bootstrap";
import React from "react";

const TabbedForm = function ({ eventKey, title, children }) {
    return (
        <Tab eventKey={eventKey} title={title}>
            {children}
        </Tab>
    )
}

const InlineForm = function ({ title, children }) {
    return (
        <Col sm={12}>
            <hr/>
            <h4 className="text-center">{title}</h4>
            <hr/>
            {children}
        </Col>
    )
}

export enum TIFormType {
    INLINE,
    TABBED
}

export interface TIFormProps {
    title: string,
    key?: string,
    type: TIFormType
}

const TIForm = function ({ title, eventKey, type, children }) {
    if(type === TIFormType.INLINE)
        return <InlineForm title={title}>{children}</InlineForm>;
    else if (type === TIFormType.TABBED)
        return <TabbedForm title={title} eventKey={eventKey}>{children}</TabbedForm>;
    else return null;
}

export {
    TIForm,
}
