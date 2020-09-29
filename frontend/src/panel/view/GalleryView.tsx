import React, {FC} from 'react';
import {
    Filter,
    TextInput,
    ReferenceInput,
    AutocompleteArrayInput, AutocompleteInput,
    TextField,
    ImageField,
    ReferenceField,
    Title,
    ListBase,
    Show,
    SimpleShowLayout,
    Create, SimpleForm, ImageInput,
    Edit,
    TopToolbar,
    CreateButton,
    ExportButton,
    SortButton,
    Pagination,
} from 'react-admin';
import { Box, Chip, useMediaQuery, Theme } from '@material-ui/core';
import GridList from "./GalleryGridList";

const GalleryTitle = ({ record }) => {
    return <span>{record.name ?? 'Upload Gallery'}</span>;
};

const GalleryFilter = (props) => (
    <Filter {...props}>
        <TextInput label="Search" source="q" alwaysOn />
        <ReferenceInput label="Album" source="albumId" reference="albums" alwaysOn>
            <AutocompleteArrayInput optionText="name" />
        </ReferenceInput>
    </Filter>
);

const ListActions: FC<any> = ({ isSmall }) => (
    <TopToolbar>
        {isSmall && <GalleryFilter context="button" />}
        <CreateButton basePath="/photos" />
        <ExportButton />
    </TopToolbar>
);

const GalleryListItem: FC<{ isSmall: boolean, customTitle: string }>
    = ({ isSmall, customTitle }) => {

    return (
        <>
            <Title defaultTitle={customTitle} />
            <ListActions isSmall={isSmall} />
            {isSmall && (
                <Box m={1}>
                    <GalleryFilter context="form" />
                </Box>
            )}
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
            <ReferenceField label="Album" source="albumId" reference="albums" target={'id'}>
                <TextField source="name" />
            </ReferenceField>
            <ImageField source="image" />
        </SimpleShowLayout>
    </Show>
)

export const GalleryEdit = props => (
    <Edit title={<GalleryTitle {...props} />} {...props}>
        <SimpleForm redirect="show">
            <TextInput disabled source="id" />
            <TextField source="title"/>
            <ReferenceInput label="Album" source="albumId" reference="albums" target={'id'} alwaysOn>
                <AutocompleteInput optionText="name" />
            </ReferenceInput>
            <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                <ImageField source="src" title="title" />
            </ImageInput>
        </SimpleForm>
    </Edit>
);

export const GalleryCreate = props => (
    <Create title={<GalleryTitle {...props} />} {...props}>
        <SimpleForm>
            <TextInput source="title" label="Title" className="d-inline" />
            <ReferenceInput label="Album" source="albumId" reference="albums" target={'id'} alwaysOn>
                <AutocompleteInput optionText="name" />
            </ReferenceInput>
            <ImageInput source="image" label="Image" accept="image/*" maxSize={500000}>
                <ImageField source="src" title="title" />
            </ImageInput>
        </SimpleForm>
    </Create>
);