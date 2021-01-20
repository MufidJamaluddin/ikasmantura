import * as React from "react";

import {
    AutocompleteArrayInput,
    AutocompleteInput,
    Create,
    Datagrid,
    DateInput,
    Edit,
    EditButton,
    Filter,
    ImageField,
    ImageInput,
    List,
    ReferenceField,
    ReferenceInput,
    RichTextField,
    Show,
    ShowButton,
    SimpleForm,
    SimpleList,
    SimpleShowLayout,
    TextField,
    TextInput,
    required, useGetIdentity
} from 'react-admin';

// @ts-ignore
import {useMediaQuery} from '@material-ui/core';

import RichTextInput from 'ra-input-rich-text';
import {dateParser} from "../../utils/DateUtil";
import {ToFormData} from "../../utils/Form";
import {FormWithImage} from "../component/FormWithImage";

const PostTitle = ({ record }) => {
    return <span>Post {record ? `"${record.title}"` : ''}</span>;
};

const PostFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn />
        <DateInput label="Dari" source={"createdAt_gte"} parse={dateParser} allowEmpty />
        <DateInput label="Sampai" source={"createdAt_lte"} parse={dateParser} allowEmpty />
        <ReferenceInput label="Penulis" source="userId" reference="users" allowEmpty>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
        <ReferenceInput label="Topik" source="topicId" reference="article_topics" allowEmpty>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

function AuthorEditButton(props) {
    if(props.userId === props.record.createdBy || props.isAdmin) {
        return (<EditButton {...props} label="Edit"/>)
    }
    return null
}

export const PostList = ({permissions, ...props}) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const { identity, loading: identityLoading } = useGetIdentity();
    const isAdmin = permissions === 'admin'

    if(identityLoading) {
        return <>Loading...</>;
    }

    return (
        <List title={props.options?.label}
              bulkActionButtons={isAdmin ? props.bulkActionButtons : false}
              filters={<PostFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={record => record.title}
                    secondaryText={record =>  <ReferenceField
                        label="Penulis" source="userId" basePath="userId" reference="users" record={record}>
                        <TextField source="name" />
                    </ReferenceField>}
                />
            ) : (
                <Datagrid rowClick="show">
                    <TextField source="id" label="ID" />
                    <TextField source="title" label="Judul" />
                    <ReferenceField label="Topik" source="topicId" reference="article_topics">
                        <TextField source="name" />
                    </ReferenceField>
                    <ImageField source="thumbnail" label="Gambar" />
                    <ReferenceField label="Penulis" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <ShowButton label="Lihat" />
                    <AuthorEditButton
                        userId={identity.id}
                        isAdmin={permissions === 'admin'}
                        {...props}
                    />
                </Datagrid>
            )}
        </List>
    )
};

export const PostShow = (props) => (
    <Show title={<PostTitle {...props} />} {...props}>
        <SimpleShowLayout>
            <TextField source="id" label="ID" />
            <ReferenceField label="Penulis" source="userId" reference="users">
                <TextField source="name" className="d-inline" />
            </ReferenceField>
            <ReferenceField label="Topik" source="topicId" reference="article_topics" >
                <TextField source="name" />
            </ReferenceField>
            <ImageField source="image" label="Gambar" />
            <TextField source="title" label="Judul" className="d-inline" />
            <RichTextField source="body" label="Isi" className="d-inline" />
        </SimpleShowLayout>
    </Show>
);

const transformData = (image, data) => {
    if(image) {
        const formData = ToFormData(data)
        formData.append('image', image)
        return formData
    }
    return data
}

export class PostEdit extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <PostTitle {...props} />
        // @ts-ignore
        return (
            <Edit transform={this.transform} title={title} {...props}>
                <SimpleForm redirect="show" encType="multipart/form-data">
                    <TextInput disabled source="id" label="ID"/>
                    <TextInput source="title" validate={[required()]} label="Judul" />
                    <ReferenceInput label="Topik" source="topicId"
                                    reference="article_topics" validate={[required()]}>
                        <AutocompleteInput optionText="name"/>
                    </ReferenceInput>
                    <ImageInput source="image" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <RichTextInput source="body" label="Isi" className="d-inline" validate={[required()]}/>
                </SimpleForm>
            </Edit>
        )
    }
}

export class PostCreate extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <PostTitle {...props} />
        // @ts-ignore
        return (
            <Create transform={this.transform} {...props}>
                <SimpleForm encType="multipart/form-data" transform={null}>
                    <TextInput source="title" label="Judul" validate={[required()]}/>
                    <ReferenceInput label="Topik" source="topicId"
                                    reference="article_topics" validate={[required()]}>
                        <AutocompleteInput optionText="name"/>
                    </ReferenceInput>
                    <ImageInput source="image" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <RichTextInput source="body" label="Isi" className="d-inline" validate={[required()]}/>
                </SimpleForm>
            </Create>
        )
    }
}
