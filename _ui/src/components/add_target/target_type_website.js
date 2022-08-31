import { TextField, FormControl, InputLabel, Select, MenuItem, Grid } from "@mui/material";

export default function TargetTypeWebsite(props) {
    return <>
        <Grid
            container
            direction="row"
            justifyContent="flex-start"
            alignItems="flex-start"
            sx={{ mb: 5 }}
        >
            <FormControl sx={{ minWidth: 110 }}>
                <InputLabel id="protocol-interval-label">Protocol</InputLabel>
                <Select
                    labelId="protocol-interval-label"
                    id="protocol-interval"
                    name='protocol'
                    value={props.values.protocol}
                    label="Protocol"
                    onChange={props.changeValues}
                >
                    <MenuItem value={'http'}>http://</MenuItem>
                    <MenuItem value={'https'}>https://</MenuItem>
                </Select>
            </FormControl>
            <TextField
                id='address'
                label='Address'
                name='address'
                value={props.values.address}
                onChange={props.changeValues}
                variant='outlined'
                sx={{ width: `calc(100% - 110px)` }}
                helperText="Include path also if you are targeting the same."
            />
        </Grid>
    </>
}