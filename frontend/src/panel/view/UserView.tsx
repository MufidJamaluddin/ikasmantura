import * as React from "react";
import {
    AutocompleteArrayInput,
    Create,
    Datagrid,
    Edit,
    EditButton,
    EmailField,
    Filter,
    List,
    PasswordInput,
    ReferenceArrayInput,
    ReferenceManyField,
    SelectInput,
    Show,
    ShowButton,
    SimpleForm,
    SimpleList,
    SimpleShowLayout,
    TextField,
    TextInput
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';

const UserTitle = ({ record }) => {
    return <span>{record.name ?? 'Create User'}</span>;
};

const UserFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
        <TextInput label="Username" source="username" allowEmpty />
        <TextInput label="Name" source="name" allowEmpty />
        <TextInput label="Email" source="email" allowEmpty />
        <SelectInput label="Address City" source="address.city" allowEmpty />
    </Filter>
);

export const UserList = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    return (
        <List title={props.options?.label} filters={<UserFilter {...props} />} {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={ record => record.name }
                    secondaryText={ record => record.username }
                    tertiaryText={ record => record.phone }
                />
                ) : (
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="name"/>
                    <TextField source="username"/>
                    <EmailField source="email"/>
                    <TextField source="phone"/>
                    <ShowButton/>
                    <EditButton/>
                </Datagrid>
                )
            }
        </List>
    )
};

export const UserView = props => (
    <Show title={<UserTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" className={"d-inline"} />
            <TextField source="username" className={"d-inline"} />
            <TextField source="name" className={"d-inline"} />
            <TextField source="email" className={"d-inline"} />
            <TextField source="phone" className={"d-inline"}/>
            <TextField source="address.street" className={"d-inline"}/>
            <TextField source="address.suite" className={"d-inline"}/>
            <TextField source="address.city" className={"d-inline"}/>
            <TextField source="address.zipcode" />
            <ReferenceManyField source="roles" reference="departments" target={"id"}>
                <Datagrid>
                    <TextField source="id" />
                    <TextField source="name" />
                    <ShowButton />
                </Datagrid>
            </ReferenceManyField>
        </SimpleShowLayout>
    </Show>
)

export const UserEdit = props => (
    <Edit title={<UserTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="name"/>
            <TextInput source="username"/>
            <TextInput source="email"/>
            <TextInput source="phone"/>
            <TextInput source="address.street" />
            <TextInput source="address.suite" />
            <TextInput source="address.city" />
            <TextInput source="address.zipcode" />
            <ReferenceArrayInput label="Roles" source="roles" reference="departments" target={"id"}>
                <AutocompleteArrayInput optionText={"name"} />
            </ReferenceArrayInput>
        </SimpleForm>
    </Edit>
)

const userCreateValidator = (values) => {
    const errors = {}
    if(values.password !== values.password_verify)
    {
        errors['password_verify'] = 'The password must match!'
    }
    return errors
}

export const UserCreate = props => (
    <Create title={<UserTitle {...props} />} {...props}>
        <SimpleForm validate={userCreateValidator}>
            <TextInput label={'Username'} source="username"/>
            <TextInput label={'Email'} source="email"/>
            <PasswordInput label={'Password'} source="password"/>
            <PasswordInput label={'Verify Password'} source="password_verify"/>
            <TextInput source="name"/>
            <TextInput source="phone"/>
            <TextInput source="address.street" />
            <TextInput source="address.suite" />
            <TextInput source="address.city" />
            <TextInput source="address.zipcode" />
            <ReferenceArrayInput label="Roles" source="roles" reference="departments" target={"id"}>
                <AutocompleteArrayInput optionText={"name"} />
            </ReferenceArrayInput>
        </SimpleForm>
    </Create>
)
