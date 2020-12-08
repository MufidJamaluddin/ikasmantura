import React, {FC} from 'react';
import {
    AutocompleteArrayInput,
    AutocompleteInput,
    Create,
    CreateButton,
    Edit,
    ExportButton,
    Filter,
    ImageField,
    ImageInput,
    ListBase,
    Pagination,
    ReferenceField,
    ReferenceInput,
    Show,
    SimpleForm,
    SimpleShowLayout,
    TextField,
    TextInput,
    Title,
    TopToolbar,
    required
} from 'react-admin';
import {Box, Theme, useMediaQuery} from '@material-ui/core';
import GridList from "./GalleryGridList";

const GalleryTitle = ({ record }) => {
    return <span>{record.name ?? 'Upload Gallery'}</span>;
};

const GalleryFilter = (props) => (
    <Filter {...props}>
        <ReferenceInput label="Album" source="albumId" reference="albums" alwaysOn>
            <AutocompleteArrayInput optionText="title" />
        </ReferenceInput>
        <TextInput label="Search" source="q" alwaysOn />
    </Filter>
);

const ListActions: FC<any> = () => (
    <>
        <Box justifyContent="flex-start">
            <GalleryFilter context="form" />
        </Box>
        <Box justifyContent="flex-end">
            <CreateButton basePath="/photos" />
            <ExportButton />
        </Box>
    </>
);

const GalleryListItem: FC<{ isSmall: boolean, customTitle: string }>
    = ({ isSmall, customTitle }) => {

    return (
        <>
            <Title defaultTitle={customTitle} />
            <TopToolbar>
            {
                isSmall ? (
                    <Box m={1}>
                        <GalleryFilter context="form" />
                    </Box>
                ) : (
                    <ListActions/>
                )
            }
            </TopToolbar>
            <Box display="flex">
                <Box width={isSmall ? 'auto' : '100%'}>
                    <GridList />
                    <Pagination rowsPerPageOptions={[10, 20, 40]} />
                </Box>
            </Box>
        </>
    )
}

export const GalleryList = props => {
    const isSmall = useMediaQuery<Theme>(theme => theme.breakpoints.down('sm'));
    return (
        <ListBase
            filters={isSmall ? <GalleryFilter {...props}/> : null}
            perPage={20}
            {...props}
        >
            <GalleryListItem customTitle={props.options?.label ?? 'Gallery'} isSmall={isSmall}/>
        </ListBase>
    )
}

export const GalleryView = props => (
    <Show title={<GalleryTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id"/>
            <TextField source="title"/>
            <ReferenceField label="Album" source="albumId" reference="albums">
                <TextField source="name" />
            </ReferenceField>
            <ImageField source="image" />
        </SimpleShowLayout>
    </Show>
)

export const GalleryEdit = props => (
    <Edit title={<GalleryTitle {...props} />} {...props}>
        <SimpleForm redirect="show" encType="multipart/form-data">
            <TextInput disabled source="id" />
            <TextInput label="Title" source="title" validate={[required()]}/>
            <ReferenceInput label="Album" source="albumId" reference="albums" validate={[required()]}>
                <AutocompleteInput optionText="title" />
            </ReferenceInput>
            <ImageInput source="image" label="Image (JPG)" accept="image/jpeg" maxSize={500000} validate={[required()]}>
                <ImageField source="src" title="title" />
            </ImageInput>
        </SimpleForm>
    </Edit>
);

export const GalleryCreate = props => (
    <Create title={<GalleryTitle {...props} />} {...props}>
        <SimpleForm encType="multipart/form-data">
            <TextInput source="title" label="Title" validate={[required()]}/>
            <ReferenceInput label="Album" source="albumId" reference="albums" validate={[required()]}>
                <AutocompleteInput optionText="title" />
            </ReferenceInput>
            <ImageInput source="image" label="Image (JPG)" accept="image/jpeg" maxSize={500000} validate={[required()]}>
                <ImageField source="src" title="title" />
            </ImageInput>
        </SimpleForm>
    </Create>
);
