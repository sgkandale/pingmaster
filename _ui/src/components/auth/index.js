import { useState } from "react"
import {
	Box, TextField, Grid, IconButton, OutlinedInput, InputLabel,
	InputAdornment, FormControl, Button, Typography
} from "@mui/material";
import { Visibility, VisibilityOff } from '@mui/icons-material';

export default function Auth() {
	const [view, setView] = useState('Login')

	const [values, setValues] = useState({
		name: '',
		password: '',
		showPassword: false,
	});


	const handleChange = (prop) => (event) => {
		setValues({ ...values, [prop]: event.target.value });
	};

	const handleClickShowPassword = () => {
		setValues({
			...values,
			showPassword: !values.showPassword,
		});
	};

	const handleMouseDownPassword = (event) => {
		event.preventDefault();
	};

	const handleFormSubmit = (event) => {
		event.preventDefault();
	}

	const renderView = () => {
		return <Grid
			container
			direction="column"
			justifyContent="center"
			alignItems="center"
			sx={{
				padding: '20px'
			}}
			component="form"
			onSubmit={handleFormSubmit}
		>
			<TextField
				id="outlined-basic"
				label="Name"
				variant="outlined"
				focused
				color="text"
				fullWidth
				value={values.name}
				onChange={handleChange('name')}
				sx={{
					marginBottom: '20px',
				}}
			/>
			<FormControl
				fullWidth
				variant="outlined"
				color="text"
			>
				<InputLabel
					htmlFor="outlined-adornment-password"
				>
					Password
				</InputLabel>
				<OutlinedInput
					id="outlined-adornment-password"
					type={values.showPassword ? 'text' : 'password'}
					value={values.password}
					onChange={handleChange('password')}
					fullWidth
					endAdornment={
						<InputAdornment position="end">
							<IconButton
								aria-label="toggle password visibility"
								onClick={handleClickShowPassword}
								onMouseDown={handleMouseDownPassword}
								edge="end"
							>
								{values.showPassword ? <VisibilityOff /> : <Visibility />}
							</IconButton>
						</InputAdornment>
					}
					label="Password"
				/>
			</FormControl>
			<Grid
				container
				direction="row"
				justifyContent="space-between"
				alignItems="center"
				sx={{
					marginTop: '40px'
				}}
			>
				<Button
					variant="outlined"
					onClick={() => {
						if (view === "Login") {
							setView("Register")
						} else {
							setView("Login")
						}
					}}
				>
					{view === "Login" ? "Register" : "Login"}
				</Button>
				<Button
					variant="contained"
					type="submit"
					onSubmit={handleFormSubmit}
				>
					{view}
				</Button>
			</Grid>
		</Grid>
	}

	return <Grid
		container
		direction="column"
		justifyContent="center"
		alignItems="center"
		sx={{
			height: '100vh'
		}}
	>
		<Box
			sx={{
				border: '1px solid',
				borderColor: 'text.disabled',
				borderRadius: '5px',
				minWidth: '400px',
				maxWidth: '600px',
				width: '100%'
			}}
		>
			<Typography
				variant="h5"
				align="center"
				sx={{
					marginTop: '10px',
					marginBottom: '10px'
				}}
			>
				{view}
			</Typography>
			{renderView()}
		</Box>
	</Grid>
}