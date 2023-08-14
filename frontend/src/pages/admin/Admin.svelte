<script>
  import Home from "../admin/Home.svelte";
  import dayjs from "dayjs";
  let listAdmin = [];
  let listAdminrule = [];
  let record = "";
  let totalrecord = 0;
  export let table_header_font = "";
  export let table_body_font = "";
  let token = localStorage.getItem("token");
  let akses_page = false;

  async function initapp() {
    const res = await fetch("/api/valid", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        page: "ADMIN-VIEW",
      }),
    });
    const json = await res.json();
    if (json.status === 400) {
      logout();
    } else if (json.status == 403) {
      alert(json.message);
    } else {
      akses_page = true;
      initAdmin();
    }
  }
  async function initAdmin() {
    const res = await fetch("/api/alladmin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({}),
    });
    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      totalrecord = record.length;
      let recordlistrule = json.listrule;
      let no = 0;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          no = no + 1;
         
          listAdmin = [
            ...listAdmin,
            {
              home_no: no,
              home_id: record[i]["admin_id"],
              home_idrule: record[i]["admin_idrule"],
              home_tipe: record[i]["admin_tipe"],
              home_nmrule: record[i]["admin_nmrule"],
              home_username: record[i]["admin_username"],
              home_nama: record[i]["admin_nama"],
              home_phone1: record[i]["admin_phone1"],
              home_phone2: record[i]["admin_phone2"],
              home_lastlogin: record[i]["admin_lastlogin"],
              home_lastipaddres: record[i]["admin_lastipaddres"],
              home_status: record[i]["admin_status"],
              home_statuscss: record[i]["admin_status_css"],
            },
          ];
        }
      }
      if (recordlistrule != null) {
        for (let i = 0; i < recordlistrule.length; i++) {
          listAdminrule = [
            ...listAdminrule,
            {
              adminrule_idruleadmin: recordlistrule[i]["adminrule_idruleadmin"],
              adminrule_nmruleadmin: recordlistrule[i]["adminrule_nmruleadmin"],
            },
          ];
        }
      }
    } else {
      logout();
    }
  }
  async function logout() {
    localStorage.clear();
    window.location.href = "/";
  }

  const handleRefreshData = (e) => {
    listAdmin = [];
    listAdminrule = [];
    totalrecord = 0;
    setTimeout(function () {
      initAdmin();
    }, 500);
  };
  initapp();
</script>

{#if akses_page == true}
  <Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listAdmin}
    {listAdminrule}
    {totalrecord}
  />
{/if}
