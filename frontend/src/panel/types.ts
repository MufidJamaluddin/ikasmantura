import {ListControllerProps} from 'ra-core/esm/controller/useListController'

export interface DatagridProps<Record>
    extends Partial<ListControllerProps<any>> {
    hasBulkActions?: boolean;
}
