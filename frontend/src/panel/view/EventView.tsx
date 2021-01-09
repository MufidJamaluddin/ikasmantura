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
        <TextInput label="Search" source="q" alwaysOn/>
        <DateInput label="From" source="createdAt_gte" parse={dateParser} alwaysOn/>
        <DateInput label="To" source="createdAt_lte" parse={dateParser} alwaysOn/>
        <TextInput label="Organizer" source="organizer" allowEmpty/>
        <TextInput label="Title" source="title" allowEmpty/>
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
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="title"/>
                    <TextField source="organizer"/>
                    <ReferenceField label="User" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <DateField source="start"/>
                    <DateField source="end"/>
                    <ImageField source="thumbnail"/>
                    <ShowButton/>
                    { permissions === 'admin' ? <EditButton/> : null }
                </Datagrid>
            )
            }
        </List>
    )
};

export const EventView = props => (
    <Show title={<EventTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" />
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

const EventValidation = (data) => {
    const errors:any = {}
    if(moment(data.start) > moment(data.end))
    {
        errors.start = 'start times must before end times!'
        errors.end = 'end times must after start times!'
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
                    <TextInput source="title" validate={[required()]}/>
                    <TextInput source="organizer" validate={[required()]}/>
                    <ReferenceInput disabled label="User" source="userId" reference="users">
                        <TextField source="name"/>
                    </ReferenceInput>
                    <RichTextInput source="description" validate={[required()]}/>
                    <ImageInput source="image" label="Image (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <DateInput source="start" validate={[required()]}/>
                    <DateInput source="end" validate={[required()]}/>
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
                    <TextInput source="title" validate={[required()]} />
                    <TextInput source="organizer" validate={[required()]} />
                    <RichTextInput source="description" validate={[required()]} />
                    <ImageInput source="image" label="Image (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                    <DateInput source="start" validate={[required()]}/>
                    <DateInput source="end" validate={[required()]}/>
                </SimpleForm>
            </Create>
        )
    }
}
