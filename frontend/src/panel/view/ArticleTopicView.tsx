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
        <TextInput label="Cari" source="q" alwaysOn />
    </Filter>
);

export const TopicList = ({ permissions, ...props}) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const isAdmin = permissions === 'admin'
    return (
        <List title={props.options?.label}
              bulkActionButtons={isAdmin ? props.bulkActionButtons : false}
              filters={<TopicFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList primaryText={record => record.name}
                />
            ) : (
                <Datagrid>
                    <TextField source="id" label="ID" />
                    <TextField source="name" label="Nama" />
                    <TextField source="icon" label="Simbol" />
                    <TextField source="description" label="Deskripsi" />
                    { permissions === 'admin' ? <EditButton label="Edit"/> : null }
                </Datagrid>
            )}
        </List>
    )
};

export const TopicEdit = props => (
    <Edit title={<TopicTitle {...props} />} {...props}>
        <SimpleForm redirect="list">
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" className="d-inline" validate={[required()]} label="Nama" />
            <TextInput source="icon" className="d-inline" validate={[required()]} label="Simbol" />
            <RichTextInput source="description" className="d-inline" validate={[required()]} label="Deskripsi" />
        </SimpleForm>
    </Edit>
);

export const TopicCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" className="d-inline" validate={[required()]} label="Nama" />
            <TextInput source="icon" className="d-inline" validate={[required()]} label="Simbol" />
            <RichTextInput source="description" className="d-inline" validate={[required()]} label="Deskripsi" />
        </SimpleForm>
    </Create>
);
