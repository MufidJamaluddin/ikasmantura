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
    required
} from 'react-admin';

import {useMediaQuery} from '@material-ui/core';

const ClassroomTitle = ({ record }) => {
    return <span>{record.name ?? 'Create Classroom'}</span>;
};

const ClassroomFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
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
                    primaryText={ record => `${record.level} - ${record.major} - ${record.seq}` }
                />
            ) : (
                <Datagrid>
                    <TextField source="id"/>
                    <TextField source="level"/>
                    <TextField source="major"/>
                    <TextField source="seq"/>
                    <ShowButton/>
                    { isAdmin ? <EditButton/> : null }
                </Datagrid>
            )
            }
        </List>
    )
}

export const ClassroomEdit = props => (
    <Edit title={<ClassroomTitle {...props} />} {...props}>
        <SimpleForm redirect="show" onSubmit={props.onSubmit}>
            <TextInput disabled source="id" />
            <TextInput source="level" validate={[required()]} />
            <TextInput source="major" validate={[required()]} />
            <TextInput source="seq" validate={[required()]} />
        </SimpleForm>
    </Edit>
)

export const ClassroomCreate = props => (
    <Create title={<ClassroomTitle {...props} />} {...props}>
        <SimpleForm onSubmit={props.onSubmit}>
            <TextInput source="level" validate={[required()]} />
            <TextInput source="major" validate={[required()]} />
            <TextInput source="seq" validate={[required()]} />
        </SimpleForm>
    </Create>
)
