import * as React from "react";
import {
    Create,
    Datagrid,
    Edit,
    EditButton,
    Filter,
    List,
    ShowButton,
    SimpleForm,
    SimpleList,
    TextField,
    TextInput,
    NumberInput,
    required
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';
import {getClassroomName} from "../Util";

const ClassroomTitle = ({ record }) => {
    return <span>{ record ? getClassroomName(record) : 'Tambah Kelas'}</span>;
};

const ClassroomFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Cari" source="q" alwaysOn />
    </Filter>
);

export const ClassroomList = ({ permissions, ...props}) => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const isAdmin = permissions === 'admin'
    return (
        <List title={props.options?.label}
              bulkActionButtons={isAdmin ? props.bulkActionButtons : false}
              filters={<ClassroomFilter {...props} />}
              {...props}>
            {isSmall ? (
                <SimpleList
                    primaryText={ record => {
                        if(record.major) return `${record.level}-${record.major}-${record.seq}`
                        else return `${record.level}-${record.seq}`
                    }}
                />
            ) : (
                <Datagrid>
                    <TextField source="id" label="ID"/>
                    <TextField source="level" label="Tingkat"/>
                    <TextField source="major" label="Jurusan"/>
                    <TextField source="seq" label="Urutan"/>
                    <ShowButton label="Lihat"/>
                    { isAdmin ? <EditButton label="Edit"/> : null }
                </Datagrid>
            )
            }
        </List>
    )
}

export const ClassroomEdit = props => (
    <Edit title={<ClassroomTitle {...props} />} {...props}>
        <SimpleForm redirect="show" onSubmit={props.onSubmit}>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="level" validate={[required()]} label="Tingkat" />
            <TextInput source="major" label="Jurusan" />
            <NumberInput source="seq" validate={[required()]} step={1} label="Urutan" />
        </SimpleForm>
    </Edit>
)

export const ClassroomCreate = props => (
    <Create title={<ClassroomTitle {...props} />} {...props}>
        <SimpleForm onSubmit={props.onSubmit}>
            <TextInput source="level" validate={[required()]} label="Tingkat" />
            <TextInput source="major" label="Jurusan" />
            <NumberInput source="seq" validate={[required()]} step={1} label="Urutan" />
        </SimpleForm>
    </Create>
)
