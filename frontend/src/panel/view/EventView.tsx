import * as React from "react";
import {
    List, Datagrid, TextField, EditButton, ShowButton,
    Filter, TextInput, DateInput, SimpleList, DateField, RichTextField,
    SimpleShowLayout, Show, ImageField, ImageInput,
    Edit, SimpleForm, ReferenceField,
    Create,
    ReferenceInput
} from 'react-admin';

import RichTextInput from 'ra-input-rich-text';

import { useMediaQuery } from '@material-ui/core';
import moment from "moment";
import {dateParser} from "../../utils/DateUtil";

const EventTitle = ({ record }) => {
    return <span>{record.name ?? 'Create Event'}</span>;
};

const EventFilter = (props) => {
    return (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn/>
        <DateInput label="From" source="createdAt_gte" parse={dateParser} alwaysOn/>
        <DateInput label="To" source="createdAt_lte" parse={dateParser} alwaysOn/>
        <TextInput label="Organizer" source="organizer" allowEmpty/>
        <TextInput label="Title" source="title" allowEmpty/>
    </Filter>)
};

export const EventList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<EventFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList rowClick={"show"}
                            primaryText={ record => record.title }
                            secondaryText={ record => record.organizer }
                            tertiaryText={ record => `${record.start} - ${record.end}` }
                />
            ) : (
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="title"/>
                    <TextField source="organizer"/>
                    <ReferenceField label="User" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <DateField source="start"/>
                    <DateField source="end"/>
                    <ShowButton/>
                    <EditButton/>
                </Datagrid>
            )
            }
        </List>
    )
};

export const EventView = props => (
    <Show title={<EventTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField disabled source="id" />
            <TextField source="title"/>
            <TextField source="organizer"/>
            <ReferenceField label="User" source="userId" reference="users">
                <TextField source="name" />
            </ReferenceField>
            <RichTextField source="description"/>
            <ImageField source="image"/>
            <DateField source="start"/>
            <DateField source="end"/>
        </SimpleShowLayout>
    </Show>
)

export const EventEdit = props => {

    const transform = data => ({
        ...data,
        start:  moment(data.start).format("YYYY-MM-DDTHH:mm:ssZ"),
        end:  moment(data.start).format("YYYY-MM-DDTHH:mm:ssZ"),
    })

    return (
        <Edit transform={transform} title={<EventTitle {...props} />} {...props}>
            <SimpleForm redirect="show">
                <TextInput disabled source="id"/>
                <TextInput source="title"/>
                <TextInput source="organizer"/>
                <ReferenceInput disabled label="User" source="userId" reference="users">
                    <TextField source="name"/>
                </ReferenceInput>
                <RichTextInput source="description"/>
                <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                    <ImageField source="src" title="title"/>
                </ImageInput>
                <DateInput source="start"/>
                <DateInput source="end"/>
            </SimpleForm>
        </Edit>
    )
}

export const EventCreate = props => {

    const transform = data => ({
        ...data,
        start:  moment(data.start).format("YYYY-MM-DDTHH:mm:ssZ"),
        end:  moment(data.start).format("YYYY-MM-DDTHH:mm:ssZ"),
    })

    return (
        <Create transform={transform} title={<EventTitle {...props} />} {...props}>
            <SimpleForm>
                <TextInput disabled source="id"/>
                <TextInput source="title"/>
                <TextInput source="organizer"/>
                <RichTextInput source="description"/>
                <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                    <ImageField source="src" title="title"/>
                </ImageInput>
                <DateInput source="start"/>
                <DateInput source="end"/>
            </SimpleForm>
        </Create>
    )
}