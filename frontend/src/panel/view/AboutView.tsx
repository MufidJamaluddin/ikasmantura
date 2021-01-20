import * as React from "react";

import {Edit, RichTextField, Show, SimpleForm, SimpleShowLayout, TextInput, required, TextField} from 'react-admin';

import RichTextInput from 'ra-input-rich-text';

const AboutTitle = () => {
    return <span>Tentang Kami</span>;
};

export const AboutShow = (props) => (
    <Show title={<AboutTitle />} {...props}>
        <SimpleShowLayout>
            <TextField source="title" label="Judul" />
            <RichTextField source="description" label="Deskripsi"  />
            <RichTextField source="vision" label="Visi" />
            <RichTextField source="mission" label="Misi" />
            <TextField source="email" label="Email" />
            <TextField source="facebook" label="Facebook" />
            <TextField source="twitter" label="Twitter" />
            <TextField source="instagram" label="Instagram" />
        </SimpleShowLayout>
    </Show>
);

export const AboutEdit = props => (
    <Edit title={<AboutTitle />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="title" validate={[required()]} label="Judul" />
            <RichTextInput source="description" validate={[required()]} label="Deskripsi" />
            <RichTextInput source="vision" validate={[required()]} label="Visi" />
            <RichTextInput source="mission" validate={[required()]} label="Misi" />
            <TextInput source="email" validate={[required()]} label="Email" />
            <TextInput source="facebook" validate={[required()]} label="Facebook" />
            <TextInput source="twitter" validate={[required()]} label="Twitter" />
            <TextInput source="instagram" validate={[required()]} label="Instagram" />
        </SimpleForm>
    </Edit>
);
