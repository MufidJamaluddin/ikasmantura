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
import {ToFormData} from "../../utils/Form";
import {FormWithImage} from "../component/FormWithImage";

const GalleryTitle = ({ record }) => {
    return <span>{record.name ?? 'Upload Gallery'}</span>;
};

const GalleryFilter = (props) => (
    <Filter {...props}>
        <ReferenceInput label="Album" source="albumId" reference="albums" alwaysOn>
            <AutocompleteArrayInput optionText="title" />
        </ReferenceInput>
        <TextInput label="Cari" source="q" alwaysOn />
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

const GalleryListItem: FC<{ permissions, isSmall: boolean, customTitle: string }>
    = ({ permissions, isSmall, customTitle }) => {
    const isAdmin = permissions === 'admin'
    return (
        <>
            <Title defaultTitle={customTitle} />
            <TopToolbar>
                {
                    isAdmin && (<ListActions/>)
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
            <GalleryListItem customTitle={props.options?.label ?? 'Galeri'} isSmall={isSmall} {...props}/>
        </ListBase>
    )
}

export const GalleryView = props => (
    <Show title={<GalleryTitle {...props} />} {...props}>
        <SimpleShowLayout className={"d-inline"}>
            <TextField source="id" label="ID"/>
            <TextField source="title" label="Judul"/>
            <ReferenceField label="Album" source="albumId" reference="albums">
                <TextField source="name" />
            </ReferenceField>
            <ImageField source="original" label="Gambar"/>
        </SimpleShowLayout>
    </Show>
)

const transformData = (image, data) => {

    const formData = ToFormData(data)

    if(image) {
        formData.delete('original')
        formData.append('image', image)
    }

    return formData
}

export class GalleryEdit extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <GalleryTitle {...props} />
        // @ts-ignore
        return (
            <Edit transform={this.transform} title={title} {...props}>
                <SimpleForm redirect="show" encType="multipart/form-data">
                    <TextInput disabled source="id"/>
                    <TextInput label="Judul" source="title" validate={[required()]}/>
                    <ReferenceInput label="Album" source="albumId"
                                    reference="albums" validate={[required()]}>
                        <AutocompleteInput optionText="title"/>
                    </ReferenceInput>
                    <ImageInput source="original" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000} validate={[required()]}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                </SimpleForm>
            </Edit>
        )
    }
}

export class GalleryCreate extends FormWithImage {

    constructor(props) {
        super(props);
    }

    transformData(image: any, data: any): any {
        return transformData(image, data)
    }

    render() {
        let props = this.props
        // @ts-ignore
        let title = <GalleryTitle {...props} />
        // @ts-ignore
        return (
            <Create transform={this.transform} title={title} {...props}>
                <SimpleForm encType="multipart/form-data">
                    <TextInput source="title" label="Judul" validate={[required()]}/>
                    <ReferenceInput label="Album" source="albumId"
                                    reference="albums" validate={[required()]}>
                        <AutocompleteInput optionText="title"/>
                    </ReferenceInput>
                    <ImageInput source="original" label="Gambar (JPG)"
                                onChange={this.dropImage}
                                accept="image/jpeg" maxSize={500000} validate={[required()]}>
                        <ImageField source="src" title="title"/>
                    </ImageInput>
                </SimpleForm>
            </Create>
        )
    }
}
