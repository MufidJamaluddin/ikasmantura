import * as React from "react";

import {
    Edit, SimpleForm, TextInput,
    RichTextField, Show, SimpleShowLayout,
    List, SimpleList
    // @ts-ignore
} from 'react-admin';

import RichTextInput from 'ra-input-rich-text';

const AboutTitle = props => {
    return <span>Tentang Kami</span>;
};

export const AboutShow = (props) => (
    <Show title={<AboutTitle {...props} />} {...props}>
        <SimpleShowLayout>
            <RichTextField source="description" />
            <RichTextField source="vision" />
            <RichTextField source="mission" />
        </SimpleShowLayout>
    </Show>
);

export const AboutEdit = props => (
    <Edit title={<AboutTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <RichTextInput source="description" />
            <RichTextInput source="vision" />
            <RichTextInput source="mission" />
        </SimpleForm>
    </Edit>
);