import * as React from 'react';
import {FC} from 'react';
import MuiGridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import GridListTileBar from '@material-ui/core/GridListTileBar';
import {makeStyles} from '@material-ui/core/styles';
import withWidth, {WithWidth} from '@material-ui/core/withWidth';
import {linkToRecord, useListContext} from 'react-admin';
import {Link} from 'react-router-dom';
import {DatagridProps} from '../types';
import {Breakpoint} from '@material-ui/core/styles/createBreakpoints';

const useStyles = makeStyles(theme => ({
    gridList: {
        margin: 0,
    },
    tileBar: {
        background:
            'linear-gradient(to top, rgba(0,0,0,0.8) 0%,rgba(0,0,0,0.4) 70%,rgba(0,0,0,0) 100%)',
    },
    placeholder: {
        backgroundColor: theme.palette.grey[300],
        height: '100%',
    },
}));

const getColsForWidth = (width: Breakpoint) => {
    if (width === 'xs') return 2;
    if (width === 'sm') return 3;
    if (width === 'md') return 4;
    if (width === 'lg') return 5;
    return 6;
};

const times = (nbChildren: number, fn: (key: number) => any) =>
    Array.from({ length: nbChildren }, (_, key) => fn(key));

const LoadingGridList: FC<GridProps & { nbItems?: number }> = ({
                                                                   width,
                                                                   nbItems = 20,
                                                               }) => {
    const classes = useStyles();
    return (
        <MuiGridList
            cellHeight={180}
            cols={getColsForWidth(width)}
            className={classes.gridList}
        >
            {' '}
            {times(nbItems, key => (
                <GridListTile key={key}>
                    <div className={classes.placeholder} />
                </GridListTile>
            ))}
        </MuiGridList>
    );
};

const LoadedGridList: FC<GridProps> = ({ width }) => {
    const { ids, data, basePath } = useListContext();
    const classes = useStyles();

    if (!ids || !data) return null;

    return (
        <MuiGridList
            cellHeight={180}
            cols={getColsForWidth(width)}
            className={classes.gridList}
        >
            {ids.map((id: string) => (
                <GridListTile
                    // @ts-ignore
                    component={Link}
                    key={id}
                    to={linkToRecord(basePath, data[id].id)}
                >
                    <img src={data[id].thumbnail} alt={data[id].title} width={'100%'} />
                    <GridListTileBar
                        className={classes.tileBar}
                        title={data[id].title}
                    />
                </GridListTile>
            ))}
        </MuiGridList>
    );
};

interface GridProps extends DatagridProps<any>, WithWidth {}

const GalleryGridList: FC<WithWidth> = ({ width }) => {
    const { loaded } = useListContext();
    return loaded ? (
        <LoadedGridList width={width} />
    ) : (
        <LoadingGridList width={width} />
    );
};

export default withWidth()(GalleryGridList);
