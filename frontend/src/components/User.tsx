import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { UsersInterface } from "../models/IUser";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import {   ListUsers } from "../services/HttpClientService";
function Users() {

  const [users, setUsers] = useState<UsersInterface[]>([]);

  const getUsers = async () => {
    let res = await ListUsers();
    if (res) {
      setUsers(res);
    }
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 75 },
    { field: "Name", headerName: "ชื่อ", width: 200 },
    { field: "Email", headerName: "อีเมล", width: 200 },
    {
        field: "Gender",
        headerName: "เพศ",
        width: 120,
        valueFormatter: (params) => params.value.Name,
    },
    {
        field: "Role",
        headerName: "บทบาท",
        width: 120,
        valueFormatter: (params) => params.value.Name,
    },
    {
        field: "Educational_background",
        headerName: "ตำแหน่ง",
        width: 125,
        valueFormatter: (params) => params.value.Name,
    },
  ];

  useEffect(() => {
    getUsers();
  }, []);

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลผู้ใช้งาน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/user/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={users}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Users;