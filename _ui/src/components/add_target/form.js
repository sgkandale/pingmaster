import { Box, TextField } from "@mui/material";
import Buttons from "./buttons";
import PingValues from "./ping_values";
import TargetType from "./target_type";

export default function Form(props) {
    return <Box
        sx={{ width: '100%', maxWidth: 800, ml: 2 }}
        component="form"
        noValidate
        autoComplete="off"
        onSubmit={props.handleFormSubmit}
    >
        <TextField
            id='name'
            label='Name'
            name='name'
            value={props.values.name}
            onChange={props.changeValues}
            variant='outlined'
            fullWidth
            helperText='Name must be unique'
            sx={{ mb: 3 }}
        />
        <TargetType
            values={props.values}
            changeValues={props.changeValues}
        />
        {
            props.values.targettype ? <PingValues
                values={props.values}
                changeValues={props.changeValues}
            /> : <></>
        }
        <Buttons
            handleFormSubmit={props.handleFormSubmit}
            stat={props.stat}
        />
    </Box>
}