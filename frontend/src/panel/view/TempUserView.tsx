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
    TextInput,
    useNotify
} from 'react-admin';

import VerifiedUserIcon from '@material-ui/icons/VerifiedUser';

import Button from '@material-ui/core/Button';

import { Confirm } from 'react-st-modal';

import {useMediaQuery} from '@material-ui/core';
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

const UserTitle = ({ record }) => {
    return <span>{record.name ?? 'User Registration'}</span>;
};

const VerifyButton = ({ record = {name: '', id: 0} }) => {

    const notify = useNotify();

    const verifyUser = async (e) => {
        e.preventDefault()

        const result = await Confirm(
            `Apakah data ${record.name} ini valid dan akan bisa digunakan untuk masuk aplikasi ini?`,
            'Verifikasi User');

        if(result)
        {
            let dataProvider = DataProviderFactory.getDataProvider()

            dataProvider.create(`verify_user/${record.id}`, { data: {} })
                .then((data:any) => {
                    notify(`User ${data.username}, atas nama ${data.name}, suskes dibuat!`)
                })
                .catch(error => {
                    notify(`Verifikasi user error: ${error.message}`, 'warning');
                })
        }
    }

    return (
        <Button onClick={verifyUser}>
            <VerifiedUserIcon/>
        </Button>
    )
}

const UserFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
        <TextInput label="Username" source="username" allowEmpty />
        <TextInput label="Name" source="name" allowEmpty />
        <TextInput label="Email" source="email" allowEmpty />
        <SelectInput label="Address City" source="address.city" allowEmpty />
    </Filter>
);

export const TempUserList = props => {
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
                    <VerifyButton/>
                    <ShowButton/>
                    <EditButton/>
                </Datagrid>
            )
            }
        </List>
    )
};

export const TempUserView = props => (
    <Show title={<UserTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" className={"d-inline"} />
            <TextField source="username" className={"d-inline"} />
            <TextField source="name" className={"d-inline"} />
            <TextField source="email" className={"d-inline"} />
            <TextField source="phone" className={"d-inline"}/>
            <TextField source="job"/>
            <TextField source="jobDesc"/>
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

export const TempUserEdit = props => (
    <Edit title={<UserTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="name"/>
            <TextInput disabled source="username"/>
            <TextInput source="email"/>
            <TextInput source="phone"/>
            <TextInput source="job"/>
            <TextInput source="jobDesc"/>
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

export const TempUserCreate = props => (
    <Create title={<UserTitle {...props} />} {...props}>
        <SimpleForm validate={userCreateValidator}>
            <TextInput label={'Username'} source="username"/>
            <TextInput label={'Email'} source="email"/>
            <PasswordInput label={'Password'} source="password"/>
            <PasswordInput label={'Verify Password'} source="password_verify"/>
            <TextInput source="name"/>
            <TextInput source="phone"/>
            <TextInput source="job"/>
            <TextInput source="jobDesc"/>
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
