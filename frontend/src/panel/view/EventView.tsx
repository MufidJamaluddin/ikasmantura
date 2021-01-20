import * as React from "react";
import {
    Create,
    Datagrid,
    DateField,
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
    required
} from 'react-admin';

import RichTextInput from 'ra-input-rich-text';

import {useMediaQuery} from '@material-ui/core';
import moment from "moment";
import {dateParser} from "../../utils/DateUtil";
import {ToFormData} from "../../utils/Form";
import {FormWithImage} from "../component/FormWithImage";

const EventTitle = ({ record }) => {
    return <span>{record.name ?? 'Create Event'}</span>;
};

const EventFilter = (props) => {
    return (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn/>
        <DateInput label="Dari" source="createdAt_gte" parse={dateParser} alwaysOn/>
        <DateInput label="Sampai" source="createdAt_lte" parse={dateParser} alwaysOn/>
        <TextInput label="Penyelenggara" source="organizer" allowEmpty/>
        <TextInput label="Judul" source="title" allowEmpty/>
    </Filter>)
};

export const EventList = ({ permissions, ...props}) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const isAdmin = permissions === 'admin'
    return (
        <List title={props.options?.label}
              bulkActionButtons={isAdmin ? props.bulkActionButtons : false}
              filters={<EventFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList primaryText={ record => record.title }
                            secondaryText={ record => record.organizer }
                            tertiaryText={ record => `${record.start} - ${record.end}` }
                />
            ) : (
                <Datagrid rowClick="show">
                    <TextField source="id" label="ID"/>
                    <TextField source="title" label="Judul"/>
                    <TextField source="organizer" label="Penyelenggara"/>
                    <ReferenceField label="Penulis" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <DateField source="start" label="Dari"/>
                    <DateField source="end" label="Sampai"/>
                    <ImageField source="thumbnail" label="Gambar"/>
                    <ShowButton label="Lihat" />
                    { permissions === 'admin' ? <EditButton label="Edit"/> : null }
                </Datagrid>
            )
            }
        </List>
    )
};

export const EventView = props => (
    <Show title={<EventTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" label="ID" />
            <TextField source="title" label="Judul" />
            <TextField source="organizer" label="Penyelenggara" />
            <ReferenceField label="Penulis" source="userId" reference="users">
                <TextField source="name" />
            </ReferenceField>
            <RichTextField source="description" label="Deskripsi"/>
            <ImageField source="image" label="Gambar"/>
            <DateField source="start" label="Waktu Dimulai (Dari)"/>
            <DateField source="end" label="Waktu Selesai (Sampai)"/>
        </SimpleShowLayout>
    </Show>
)

const EventValidation = (data) => {
    const errors:any = {}
    if(moment(data.start) > moment(data.end))
    {
        errors.start = 'waktu dimulai harus berada sebelum waktu selesai!'
        errors.end = 'waktu selesai harus berada setelah waktu dimulai!'
    }
    return errors
}

const transformData = (image, data) => {
    let transformedData = {
        ...data,
        start:  moment(data.start).format("YYYY-MM-DDTHH:mm:ssZ"),
        end:  moment(data.end).format("YYYY-MM-DDTHH:mm:ssZ"),
    }

    if(image) {
        const formData = ToFormData(transformedData)
        formData.append('image', image)
        return formData
    }

    return transformedData
}

export class EventEdit extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <EventTitle {...props} />
        // @ts-ignore
        return (
            <Edit transform={this.transform} title={title} {...props}>
                <SimpleForm validate={EventValidation} redirect="show" encType="multipart/form-data">
                    <TextInput disabled source="id"/>
                    <TextInput source="title" validate={[required()]} label="Judul"/>
                    <TextInput source="organizer" validate={[required()]} label="Penyelenggara"/>
                    <ReferenceInput disabled label="Penulis" source="userId" reference="users">
                        <TextField source="name"/>
                    </ReferenceInput>
                    <RichTextInput source="description" label="Deskripsi" validate={[required()]}/>
                    <ImageInput source="image" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <DateInput source="start" validate={[required()]} label="Waktu Dimulai (Dari)"/>
                    <DateInput source="end" validate={[required()]} label="Waktu Selesai (Sampai)"/>
                </SimpleForm>
            </Edit>
        )
    }
}

export class EventCreate extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <EventTitle {...props} />
        // @ts-ignore
        return (
            <Create transform={this.transform}
                    title={title} {...props}>
                <SimpleForm validate={EventValidation} encType="multipart/form-data">
                    <TextInput disabled source="id"/>
                    <TextInput source="title" validate={[required()]} label="Judul" />
                    <TextInput source="organizer" validate={[required()]} label="Penyelenggara" />
                    <RichTextInput source="description" validate={[required()]} label="Deskripsi" />
                    <ImageInput source="image" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <DateInput source="start" validate={[required()]} label="Waktu Dimulai (Dari)" />
                    <DateInput source="end" validate={[required()]} label="Waktu Selesai (Sampai)" />
                </SimpleForm>
            </Create>
        )
    }
}
