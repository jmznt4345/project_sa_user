import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { UsersInterface } from "../models/IUser";
import { GenderInterface } from "../models/IGender";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import { CreateUser, GetGender, GetEducational_background, GetRole } from "../services/HttpClientService";
import React, { useEffect, useState } from "react";
import { RoleInterface } from "../models/IRole";
import { Educational_backgroundInterface } from "../models/IEducational_background";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function UserCreate() {
  
  const [user, setUser] = React.useState<UsersInterface>({});
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);
  const [gender, setGender] = React.useState<GenderInterface[]>([]);
  const [role, setRole] = React.useState<GenderInterface[]>([]);
  const [educational_background, setEducational_background] = React.useState<GenderInterface[]>([]);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof UserCreate;
    const { value } = event.target;
    setUser({ ...user, [id]: value });

  };

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Name: user.Name,
      Email: user.Name,
      Phonenumber: user.Phonenumber,
      Password: user.Password,

      Educational_backgroundID: convertType(user.Educational_backgroundID),
      RoleID: convertType(user.RoleID),
      GenderID: convertType(user.GenderID),
    };
    console.log(data);
    let res = await CreateUser(data);
    console.log(res);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof user;
    const value = event.target.value;
    setUser({
      ...user,
      [name]: value,
    });
    console.log(`[${name}]: ${value}`);
  };

  const getRole = async () => {
    let res = await GetRole();
    if (res) {
      setRole(res);
      console.log("Load Role Complete");
    }
    else {
      console.log("Load Role InComplete!!!!");
    }
  };

  const getGender = async () => {
    let res = await GetGender();
    if (res) {
      setGender(res);
      console.log("Load Gender Complete");
    }
    else {
      console.log("Load Gender InComplete!!!!");
    }
  };

  const getEducational_background = async () => {
    let res = await GetEducational_background();
    if (res) {
      setEducational_background(res);
      console.log("Load Educational_background Complete");
    }
    else {
      console.log("Load Educational_background InComplete!!!!");
    }
  };

  useEffect(() => {
    getRole();
    getGender();
    getEducational_background();
  }, []);

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>

      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>

      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              สร้างข้อมูลสมาชิก
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>

          <Grid item xs={4}>
            <p>Name</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลชื่อ"
                value={user.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>Phone Number</p>
              <TextField
                id="Phonenumber"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลโทรศัพท์มือถือ"
                value={user.Phonenumber || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>Gender</p>
              <Select
                required
                defaultValue={"0"}
                // value={request.RHD_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "GenderID",
                }}
              >
                <MenuItem value={"0"}>กรุณาเลือกเพศ</MenuItem>
                {gender?.map((item: GenderInterface) =>
                  <MenuItem
                    key={item.ID}
                    value={item.ID}
                  >
                    {item.Name}
                  </MenuItem>
                )}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Email</p>
              <TextField
                id="Email"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลอีกเมล"
                value={user.Email || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Password</p>
              <TextField
                id="Password"
                variant="outlined"
                type="password"
                size="medium"
                placeholder="กรุณากรอกข้อมูลรหัสผ่าน"
                value={user.Password || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Role</p>
              <Select
                required
                defaultValue={"0"}
                // value={request.RHD_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "RoleID",
                }}
              >
                <MenuItem value={"0"}>กรุณาเลือกบทบาท</MenuItem>
                {role?.map((item: RoleInterface) =>
                  <MenuItem
                    key={item.ID}
                    value={item.ID}
                  >
                    {item.Name}
                  </MenuItem>
                )}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Educational_background</p>
              <Select
                required
                defaultValue={"0"}
                // value={request.RHD_ID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Educational_backgroundID",
                }}
              >
                <MenuItem value={"0"}>กรุณาเลือกวุฒิการศึกษา</MenuItem>
                {educational_background?.map((item: Educational_backgroundInterface) =>
                  <MenuItem
                    key={item.ID}
                    value={item.ID}
                  >
                    {item.Name}
                  </MenuItem>
                )}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/users"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default UserCreate;