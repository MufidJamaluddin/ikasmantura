import * as React from "react";
import {
    AutocompleteInput,
    Create,
    Datagrid,
    Edit,
    EditButton,
    Filter,
    List,
    ReferenceField,
    ReferenceInput,
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

const DepartmentTitle = ({ record }) => {
    return <span>{record.name ?? 'Tambah Departemen'}</span>;
};

const DepartmentFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Departemen" source="name" alwaysOn />
    </Filter>
);

export const DepartmentList = ({ permissions, ...props}) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const isAdmin = permissions === 'admin'
    return (
        <List
            title={props.options?.label}
            bulkActionButtons={isAdmin ? props.bulkActionButtons : false}
            filters={<DepartmentFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={ record => record.name }
                />
            ) : (
                <Datagrid>
                    <TextField source="id" label="ID"/>
                    <TextField source="name" label="Nama"/>
                    <ReferenceField label="Pejabat" source="userId" reference="users">
                        <TextField source="name" />
                    </ReferenceField>
                    <ShowButton label="Lihat"/>
                    { isAdmin ? <EditButton label="Edit"/> : null }
                </Datagrid>
            )
            }
        </List>
    )
};

export const DepartmentShow = props => {
    return (
        <Show title={<DepartmentTitle {...props}/>} {...props}>
            <SimpleShowLayout>
                <TextField source="id" label="ID" className="d-inline"/>
                <TextField source="name" label="Nama" />
                <ReferenceField label="Pejabat" source="userId" reference="users">
                    <TextField source="name" />
                </ReferenceField>
            </SimpleShowLayout>
        </Show>
    )
}

export const DepartmentEdit = props => {
    return (
        <Edit title={<DepartmentTitle {...props}/>} {...props}>
            <SimpleForm className={"d-inline"}>
                <TextInput disabled source="id" label="ID" />
                <TextInput label="Nama Departemen"  source="name" validate={[required()]} />
                <ReferenceInput label="Pejabat" source="userId" reference="users" validate={[required()]}>
                    <AutocompleteInput optionText="name" className="d-inline" />
                </ReferenceInput>
            </SimpleForm>
        </Edit>
    )
}

export const DepartmentCreate = props => {
    return (
        <Create title={<DepartmentTitle {...props}/>} {...props}>
            <SimpleForm className={"d-inline"}>
                <TextInput label="Nama Departemen" source="name" validate={[required()]}/>
                <ReferenceInput label="Pejabat" source="userId" reference="users" validate={[required()]}>
                    <AutocompleteInput optionText="name" className="d-inline" />
                </ReferenceInput>
            </SimpleForm>
        </Create>
    )
}
