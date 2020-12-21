import * as React from "react";

import {Edit, RichTextField, Show, SimpleForm, SimpleShowLayout, TextInput, required} from 'react-admin';

import RichTextInput from 'ra-input-rich-text';

const AboutTitle = props => {
    return <span>Tentang Kami</span>;
};

export const AboutShow = (props) => (
    <Show title={<AboutTitle {...props} />} {...props}>
        <SimpleShowLayout>
            <RichTextField source="title" />
            <RichTextField source="description" />
            <RichTextField source="vision" />
            <RichTextField source="mission" />
            <RichTextField source="email" />
            <RichTextField source="facebook" />
            <RichTextField source="twitter" />
            <RichTextField source="instagram" />
        </SimpleShowLayout>
    </Show>
);

export const AboutEdit = props => (
    <Edit title={<AboutTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <RichTextInput source="title" validate={[required()]} />
            <RichTextInput source="description" validate={[required()]} />
            <RichTextInput source="vision" validate={[required()]} />
            <RichTextInput source="mission" validate={[required()]} />
            <RichTextInput source="email" validate={[required()]} />
            <RichTextInput source="facebook" validate={[required()]} />
            <RichTextInput source="twitter" validate={[required()]} />
            <RichTextInput source="instagram" validate={[required()]} />
        </SimpleForm>
    </Edit>
);
