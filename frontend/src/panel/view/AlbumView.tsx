import * as React from "react";
import {
    AutocompleteArrayInput,
    Create,
    Datagrid,
    Edit,
    EditButton,
    Filter,
    ImageField,
    List,
    ReferenceField,
    ReferenceInput,
    ReferenceManyField,
    Show,
    ShowButton,
    SimpleForm,
    SimpleList,
    SimpleShowLayout,
    TextField,
    TextInput,
    required
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';

const AlbumTitle = ({ record }) => {
    return <span>{record.name ?? 'Tambah Album'}</span>;
};

const AlbumFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn />
        <ReferenceInput label="Penulis" source="userId" reference="users" allowEmpty>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

export const AlbumList = ({ permissions, ...props }) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<AlbumFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={ record => record.title }
                    secondaryText={record =>  <ReferenceField
                        label="Penulis" source="userId" basePath="userId" reference="users" record={record}>
                        <TextField source="name" />
                    </ReferenceField>}
                />
            ) : (
                <Datagrid rowClick="show">
                    <TextField source="id" label="ID"/>
                    <TextField source="title" label="Judul" />
                    <ReferenceField label="Penulis" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <ShowButton label="Lihat"/>
                    { permissions === 'admin' ? <EditButton label="Edit"/> : null }
                </Datagrid>
            )
            }
        </List>
    )
};

export const AlbumView = props => (
    <Show title={<AlbumTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" label="ID" />
            <TextField source="title" label="Judul" />
            <ReferenceField label="Penulis" source="userId" reference="users">
                <TextField source="name" />
            </ReferenceField>
            <ReferenceManyField label="Galeri" source="id" reference="photos" target="albumId">
                <Datagrid>
                    <TextField source={'title'} label="Nama Foto"/>
                    <ImageField source={'thumbnail'} label="Foto" />
                </Datagrid>
            </ReferenceManyField>
        </SimpleShowLayout>
    </Show>
)

export const AlbumEdit = props => (
    <Edit title={<AlbumTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="title" validate={[required()]} label="Judul" />
            <ReferenceField label="Penulis" source="userId" reference="users">
                <TextField source="name"/>
            </ReferenceField>
        </SimpleForm>
    </Edit>
)

export const AlbumCreate = props => (
    <Create title={<AlbumTitle {...props} />} {...props}>
        <SimpleForm>
            <TextInput label="Judul" source="title" validate={[required()]} />
        </SimpleForm>
    </Create>
)
