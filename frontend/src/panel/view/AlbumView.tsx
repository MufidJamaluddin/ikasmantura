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
    return <span>{record.name ?? 'Create Album'}</span>;
};

const AlbumFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
        <ReferenceInput label="Author" source="userId" reference="users" allowEmpty>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

export const AlbumList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<AlbumFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={ record => record.title }
                    secondaryText={record =>  <ReferenceField
                        label="Author" source="userId" basePath="userId" reference="users" record={record}>
                        <TextField source="name" />
                    </ReferenceField>}
                />
            ) : (
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="title"/>
                    <ReferenceField label="Author" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <ShowButton/>
                    <EditButton/>
                </Datagrid>
            )
            }
        </List>
    )
};

export const AlbumView = props => (
    <Show title={<AlbumTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" />
            <TextField source="title"/>
            <ReferenceField label="Author" source="userId" reference="users">
                <TextField source="name" />
            </ReferenceField>
            <ReferenceManyField label={"Gallery"} source="id" reference="photos" target="albumId">
                <Datagrid>
                    <TextField source={'title'}/>
                    <ImageField source={'thumbnail'} />
                </Datagrid>
            </ReferenceManyField>
        </SimpleShowLayout>
    </Show>
)

export const AlbumEdit = props => (
    <Edit title={<AlbumTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="title" validate={[required()]} />
            <ReferenceField label="Author" source="userId" reference="users">
                <TextField source="name"/>
            </ReferenceField>
        </SimpleForm>
    </Edit>
)

export const AlbumCreate = props => (
    <Create title={<AlbumTitle {...props} />} {...props}>
        <SimpleForm>
            <TextInput label={'Title'} source="title" validate={[required()]} />
        </SimpleForm>
    </Create>
)
