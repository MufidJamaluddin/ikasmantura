import * as React from "react";
import {
    BooleanField,
    BooleanInput,
    Create,
    Datagrid,
    Edit,
    EditButton,
    EmailField,
    Filter,
    List,
    PasswordInput, ReferenceArrayField,
    ReferenceArrayInput,
    SelectArrayInput,
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
import {getClassroomName} from "../Util";

const UserTitle = ({ record }) => {
    return <span>{record.name ?? 'Registrasi Anggota'}</span>;
};

const VerifyButton = ({ record = {name: '', username: '', id: 0} }) => {

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
                    notify(`User ${record.username}, atas nama ${record.name}, suskes dibuat!`)
                })
                .catch(error => {
                    notify(`Verifikasi user error: ${error.message}`, 'warning');
                })
        }
    }

    return (
        <Button onClick={verifyUser} className="MuiButton-textPrimary">
            <VerifiedUserIcon/>
            &nbsp; Verifikasi
        </Button>
    )
}

const UserFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn />
        <TextInput label="Username" source="username" allowEmpty />
        <TextInput label="Nama" source="name" allowEmpty />
        <TextInput label="Email" source="email" allowEmpty />
        <SelectInput label="Kota Tinggal" source="address.city" allowEmpty />
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
                <Datagrid rowClick="show">
                    <TextField source="id" label="ID"/>
                    <TextField source="name" label="Nama"/>
                    <TextField source="username" label="Username"/>
                    <EmailField source="email" label="Email"/>
                    <TextField source="phone" label="HP"/>
                    <VerifyButton/>
                    <ShowButton label="Lihat"/>
                    <EditButton label="Edit"/>
                </Datagrid>
            )
            }
        </List>
    )
};

export const TempUserView = props => (
    <Show title={<UserTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" className={"d-inline"} label="ID" />
            <TextField source="username" className={"d-inline"} label="Username" />
            <TextField source="name" className={"d-inline"} label="Nama" />
            <TextField source="email" className={"d-inline"} label="Email" />
            <BooleanField source="emailValid" label="Email Valid (Telah Diverifikasi)" className={"d-inline"}/>
            <TextField source="phone" className={"d-inline"} label="HP" />
            <TextField source="forceYear" label="Tahun Lulus" />
            <TextField source="job" label="Pekerjaan"/>
            <TextField source="jobDesc" label="Deskripsi Pekerjaan" />
            <TextField source="address.street" className={"d-inline"} label="Alamat" />
            <TextField source="address.suite" className={"d-inline"} label="Kompleks"/>
            <TextField source="address.city" className={"d-inline"} label="Kota Tinggal"/>
            <TextField source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayField source="classrooms" reference="classrooms" label="Kelas">
                <Datagrid>
                    <TextField source="id" label="ID" />
                    <TextField source="level" label="Tingkat" />
                    <TextField source="major" label="Jurusan" />
                    <TextField source="seq" label="Urutan" />
                </Datagrid>
            </ReferenceArrayField>
        </SimpleShowLayout>
    </Show>
)

export const TempUserEdit = props => (
    <Edit title={<UserTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="name" label="Nama" />
            <TextInput disabled source="username" label="Username" />
            <TextInput source="email" label="Email" />
            <BooleanInput source="emailValid" label="Email Valid (Telah Diverifikasi)"/>
            <TextInput source="phone" label="HP" />
            <TextInput source="forceYear" label="Tahun Lulus" />
            <TextInput source="job" label="Pekerjaan" />
            <TextInput source="jobDesc" label="Deskripsi Pekerjaan" />
            <TextInput source="address.street" label="Alamat" />
            <TextInput source="address.suite" label="Kompleks" />
            <TextInput source="address.city" label="Kota Tinggal" />
            <TextInput source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayInput source="classrooms" reference="classrooms" label="Kelas">
                <SelectArrayInput optionText={getClassroomName} />
            </ReferenceArrayInput>
        </SimpleForm>
    </Edit>
)

const userCreateValidator = (values) => {
    const errors = {}
    if(values.password !== values.password_verify)
    {
        errors['password_verify'] = 'Password harus sama!'
    }
    return errors
}

export const TempUserCreate = props => (
    <Create title={<UserTitle {...props} />} {...props}>
        <SimpleForm validate={userCreateValidator}>
            <TextInput label={'Username'} source="username"/>
            <TextInput label={'Email'} source="email"/>
            <BooleanInput source="emailValid" label="Email Valid (Telah Diverifikasi)"/>
            <PasswordInput label={'Password'} source="password"/>
            <PasswordInput label={'Verifikasi Password'} source="password_verify"/>
            <TextInput source="name" label="Nama Lengkap"/>
            <TextInput source="phone" label="HP"/>
            <TextInput source="forceYear" label="Tahun Lulus"/>
            <TextInput source="job" label="Pekerjaan" />
            <TextInput source="jobDesc" label="Deskripsi Pekerjaan" />
            <TextInput source="address.street" label="Alamat" />
            <TextInput source="address.suite" label="Kompleks" />
            <TextInput source="address.city" label="Kota" />
            <TextInput source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayInput source="classrooms" reference="classrooms" label="Kelas">
                <SelectArrayInput optionText={getClassroomName} />
            </ReferenceArrayInput>
        </SimpleForm>
    </Create>
)
