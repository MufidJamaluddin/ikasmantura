import * as React from "react";
import {
    Create,
    Datagrid,
    Edit,
    EditButton,
    EmailField,
    BooleanField,
    Filter,
    List,
    PasswordInput,
    RadioButtonGroupInput,
    ReferenceArrayField, ReferenceArrayInput,
    ReferenceManyField,
    SelectArrayInput,
    SelectInput,
    Show,
    ShowButton,
    SimpleForm,
    SimpleList,
    SimpleShowLayout,
    TextField,
    TextInput,
    BooleanInput
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';
import {getClassroomName} from "../Util";

const UserTitle = ({ record }) => {
    return <span>{record.name ?? 'Create User'}</span>;
};

const UserFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn />
        <TextInput label="Username" source="username" allowEmpty />
        <TextInput label="Nama" source="name" allowEmpty />
        <TextInput label="Email" source="email" allowEmpty />
        <SelectInput label="Kota Tinggal" source="address.city" allowEmpty />
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
                <Datagrid rowClick="show">
                    <TextField source="id" label="ID" />
                    <TextField source="name" label="Nama" />
                    <TextField source="username" label="Username" />
                    <EmailField source="email" label="Email" />
                    <TextField source="phone" label="HP" />
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
            <TextField source="id" className={"d-inline"} label="ID" />
            <TextField source="username" className={"d-inline"} label="Username" />
            <TextField source="name" className={"d-inline"} label="Nama" />
            <TextField source="email" className={"d-inline"} label="Email" />
            <BooleanField source="emailValid" label="Email Valid (Telah Diverifikasi)"/>
            <TextField source="phone" className={"d-inline"} label="HP" />
            <TextField source="forceYear" label="Tahun Lulus" />
            <TextField source="job" label="Pekerjaan"/>
            <TextField source="jobDesc" label="Deskripsi Pekerjaan" />
            <TextField source="address.street" className={"d-inline"} label="Alamat" />
            <TextField source="address.suite" className={"d-inline"} label="Kompleks" />
            <TextField source="address.city" className={"d-inline"} label="Kota Tinggal" />
            <TextField source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayField source="classrooms" reference="classrooms" label="Kelas">
                <Datagrid>
                    <TextField source="id" label="ID" />
                    <TextField source="level" label="Tingkat" />
                    <TextField source="major" label="Jurusan" />
                    <TextField source="seq" label="Urutan" />
                </Datagrid>
            </ReferenceArrayField>
            <TextField source="role" label="Hak Akses Aplikasi/Peran" />
            <ReferenceManyField label="Jabatan" reference="departments" target="userId">
                <Datagrid>
                    <TextField source="id" label="ID"/>
                    <TextField source="name" label="Nama Jabatan"/>
                    <ShowButton label="Lihat"/>
                </Datagrid>
            </ReferenceManyField>
        </SimpleShowLayout>
    </Show>
)

export const UserEdit = props => (
    <Edit title={<UserTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextInput source="name" label="Nama"/>
            <TextInput disabled source="username"/>
            <TextInput source="email" label="Email" />
            <BooleanInput source="emailValid" label="Email Valid (Telah Diverifikasi)"/>
            <TextInput source="phone" label="HP"/>
            <TextInput source="forceYear" label="Tahun Lulus"/>
            <TextInput source="job" label="Pekerjaan"/>
            <TextInput source="jobDesc" label="Deskripsi Pekerjaan"/>
            <TextInput source="address.street" label="Alamat" />
            <TextInput source="address.suite" label="Kompleks" />
            <TextInput source="address.city" label="Kota Tinggal" />
            <TextInput source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayInput source="classrooms" reference="classrooms" label="Kelas">
                <SelectArrayInput optionText={getClassroomName} />
            </ReferenceArrayInput>
            <RadioButtonGroupInput source="role" label="Hak Akses Aplikasi/Peran" choices={[
                { id: 'member', name: 'Anggota' },
                { id: 'admin', name: 'Admin' },
            ]} />
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
            <BooleanInput source="emailValid" label="Email Valid (Telah Diverifikasi)"/>
            <PasswordInput label={'Password'} source="password"/>
            <PasswordInput label={'Verifikasi Password'} source="password_verify"/>
            <TextInput source="name" label="Nama" />
            <TextInput source="phone" label="HP" />
            <TextInput source="forceYear" label="Tahun Lulus" />
            <TextInput source="job" label="Pekerjaan" />
            <TextInput source="jobDesc" label="Deskripsi Pekerjaan" />
            <TextInput source="address.street" label="Alamat" />
            <TextInput source="address.suite" label="Kompleks" />
            <TextInput source="address.city" label="Kota" />
            <TextInput source="address.zipcode" label="Kode Pos" />
            <ReferenceArrayInput source="classrooms" reference="classrooms" label="Kelas">
                <SelectArrayInput optionText={getClassroomName} />
            </ReferenceArrayInput>
            <RadioButtonGroupInput source="role" label="Hak Akses Aplikasi/Peran" choices={[
                { id: 'member', name: 'Anggota' },
                { id: 'admin', name: 'Admin' },
            ]} />
        </SimpleForm>
    </Create>
)
