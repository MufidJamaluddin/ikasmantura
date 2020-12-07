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
    TextInput
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';

const DepartmentTitle = ({ record }) => {
    return <span>{record.name ?? 'Create Department'}</span>;
};

const DepartmentFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Department" source="name" alwaysOn />
    </Filter>
);

export const DepartmentList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<DepartmentFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList rowClick={"show"}
                            primaryText={ record => record.name }
                />
            ) : (
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="name"/>
                    <ReferenceField source="userId" reference="users">
                        <TextField optionText="name" className="d-inline" />
                    </ReferenceField>
                    <ShowButton/>
                    <EditButton/>
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
                <TextField disabled source="id" className={"d-inline"}/>
                <TextField source="name" />
                <ReferenceField source="userId" reference="users">
                    <TextField optionText="name" className="d-inline" />
                </ReferenceField>
            </SimpleShowLayout>
        </Show>
    )
}

export const DepartmentEdit = props => {
    return (
        <Edit title={<DepartmentTitle {...props}/>} {...props}>
            <SimpleForm className={"d-inline"}>
                <TextInput disabled source="id" />
                <TextInput source="name"/>
                <ReferenceInput source="userId" reference="users">
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
                <TextInput source="name"/>
                <ReferenceInput source="userId" reference="users">
                    <AutocompleteInput optionText="name" className="d-inline" />
                </ReferenceInput>
            </SimpleForm>
        </Create>
    )
}
