import {ListControllerProps} from 'ra-core/esm/controller/useListController'
import {Record} from 'ra-core'

export interface DatagridProps<RecordType = Record>
    extends Partial<ListControllerProps<RecordType>> {
    hasBulkActions?: boolean;
}
