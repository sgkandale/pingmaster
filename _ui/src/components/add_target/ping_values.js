import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";


export default function PingValues(props) {
    return <>
        <FormControl fullWidth sx={{ mb: 5 }}>
            <InputLabel id="ping-interval-label">Ping Interval</InputLabel>
            <Select
                labelId="ping-interval-label"
                id="ping-interval"
                name='pinginterval'
                value={props.values.pinginterval}
                label="Ping Interval"
                onChange={props.changeValues}
            >
                <MenuItem value={'10s'}>10 Seconds</MenuItem>
                <MenuItem value={'30s'}>30 Seconds</MenuItem>
                <MenuItem value={'1m'}>1 Minute</MenuItem>
                <MenuItem value={'2m'}>2 Minutes</MenuItem>
                <MenuItem value={'5m'}>5 Minutes</MenuItem>
                <MenuItem value={'10m'}>10 Minutes</MenuItem>
                <MenuItem value={'30m'}>30 Minutes</MenuItem>
                <MenuItem value={'1h'}>1 Hour</MenuItem>
                <MenuItem value={'2h'}>2 Hours</MenuItem>
                <MenuItem value={'6h'}>6 Hours</MenuItem>
                <MenuItem value={'24h'}>24 Hours</MenuItem>
            </Select>
        </FormControl>
        <FormControl fullWidth>
            <InputLabel id="ping-timeout-label">Ping Timeout</InputLabel>
            <Select
                labelId="ping-timeout-label"
                id="ping-timeout"
                name='pingtimeout'
                value={props.values.pingtimeout}
                label="Ping Timeout"
                onChange={props.changeValues}
            >
                <MenuItem value={'1s'}>1 Second</MenuItem>
                <MenuItem value={'5s'}>5 Seconds</MenuItem>
                <MenuItem value={'10s'}>10 Seconds</MenuItem>
                <MenuItem value={'30s'}>30 Seconds</MenuItem>
            </Select>
        </FormControl>
    </>
}