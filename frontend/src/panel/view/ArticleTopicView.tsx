import * as React from "react";

import {
    Create,
    Datagrid,
    Edit,
    EditButton,
    Filter,
    List,
    SimpleForm,
    SimpleList,
    TextField,
    TextInput,
    required
} from 'react-admin';

// @ts-ignore
import {useMediaQuery} from '@material-ui/core';

import RichTextInput from 'ra-input-rich-text';

const TopicTitle = ({ record }) => {
    return <span>Topic {record ? `"${record.title}"` : ''}</span>;
};

const TopicFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
    </Filter>
);

export const TopicList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<TopicFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList primaryText={record => record.name}
                />
            ) : (
                <Datagrid>
                    <TextField source="id" />
                    <TextField source="name" />
                    <TextField source="icon" />
                    <TextField source="description" />
                    <EditButton />
                </Datagrid>
            )}
        </List>
    )
};

export const TopicEdit = props => (
    <Edit title={<TopicTitle {...props} />} {...props}>
        <SimpleForm redirect="list">
            <TextInput disabled source="id" />
            <TextInput source="name" className="d-inline" validate={[required()]} />
            <TextInput source="icon" className="d-inline" validate={[required()]} />
            <RichTextInput source="description" className="d-inline" validate={[required()]} />
        </SimpleForm>
    </Edit>
);

export const TopicCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput disabled source="id" />
            <TextInput source="name" className="d-inline" validate={[required()]} />
            <TextInput source="icon" className="d-inline" validate={[required()]} />
            <RichTextInput source="description" className="d-inline" validate={[required()]} />
        </SimpleForm>
    </Create>
);
