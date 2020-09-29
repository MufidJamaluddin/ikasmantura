import * as React from "react";

import {
    List, Datagrid, TextField, ReferenceField, Filter, SimpleList,
    Edit, SimpleForm, TextInput, ReferenceInput, SelectInput, EditButton, ShowButton,
    Create, ImageField,
    RichTextField, Show, SimpleShowLayout, AutocompleteArrayInput, DateInput, ImageInput
    // @ts-ignore
} from 'react-admin';

// @ts-ignore
import { useMediaQuery } from '@material-ui/core';

import RichTextInput from 'ra-input-rich-text';
import {dateParser} from "../../utils/DateUtil";

const PostTitle = ({ record }) => {
    return <span>Post {record ? `"${record.title}"` : ''}</span>;
};

const PostFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
        <DateInput label={"From"} source={"createdAt_gte"} parse={dateParser} allowEmpty />
        <DateInput label={"To"} source={"createdAt_lte"} parse={dateParser} allowEmpty />
        <ReferenceInput label="User" source="userId" reference="users" allowEmpty>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

export const PostList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<PostFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList rowClick={"show"}
                    primaryText={record => record.title}
                    secondaryText={record =>  <ReferenceField
                        label="User" source="userId" basePath="userId" reference="users" record={record}>
                        <TextField source="name" />
                    </ReferenceField>}
                />
            ) : (
                <Datagrid>
                    <TextField source="id" />
                    <ReferenceField label="User" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <TextField source="title" />
                    <ShowButton />
                    <EditButton />
                </Datagrid>
            )}
        </List>
    )
};

export const PostShow = (props) => (
    <Show title={<PostTitle {...props} />} {...props}>
        <SimpleShowLayout>
            <TextField disabled source="id" />
            <ReferenceField source="userId" reference="users">
                <TextField optionText="name" className="d-inline" />
            </ReferenceField>
            <ImageField source="image" />
            <TextField source="title" className="d-inline" />
            <RichTextField source="body" className="d-inline" />
        </SimpleShowLayout>
    </Show>
);

export const PostEdit = props => (
    <Edit title={<PostTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                <ImageField source="src" title="title" />
            </ImageInput>
            <TextInput source="title" className="d-inline" />
            <RichTextInput source="body" className="d-inline" />
        </SimpleForm>
    </Edit>
);

export const PostCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <TextInput source="title" label="Title" className="d-inline" />
            <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                <ImageField source="src" title="title" />
            </ImageInput>
            <RichTextInput source="body" className="d-inline" />
        </SimpleForm>
    </Create>
);