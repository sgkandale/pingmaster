import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";
import TargetTypeWebsite from "./target_type_website";

const TargetType_Website = 'website'

export default function TargetType(props) {

    const renderTypeOptions = () => {
        if (props.values.targettype === TargetType_Website) {
            return <TargetTypeWebsite
                values={props.values}
                changeValues={props.changeValues}
            />
        }
        return <></>
    }

    return <>
        <FormControl fullWidth sx={{ mb: 5 }}>
            <InputLabel id="target-type-label">Target Type</InputLabel>
            <Select
                labelId="target-type-label"
                id="target-type"
                name='targettype'
                value={props.values.targettype}
                label="Target Type"
                onChange={props.changeValues}
            >
                <MenuItem value={TargetType_Website}>Website</MenuItem>
            </Select>
        </FormControl>
        {renderTypeOptions()}
    </>
}