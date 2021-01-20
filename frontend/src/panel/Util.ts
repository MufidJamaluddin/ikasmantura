export function getClassroomName(record) {
    if(record) {
        if (record.major) {
            return `${record.level}-${record.major}-${record.seq}`
        }
        return `${record.level}-${record.seq}`
    }
    return ''
}
