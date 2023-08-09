<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listCurr = [];
    let listBank = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "MASTER-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/master", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_curr = json.listcurr;
            let record_bank = json.listbank;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    let endmaster = ""
                    if (record[i]["master_start"] != record[i]["master_end"]){
                        endmaster = record[i]["master_end"]
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["master_id"],
                            home_idcurr: record[i]["master_idcurr"],
                            home_start: record[i]["master_start"],
                            home_end: endmaster,
                            home_name: record[i]["master_name"],
                            home_owner: record[i]["master_owner"],
                            home_phone1: record[i]["master_phone1"],
                            home_phone2: record[i]["master_phone1"],
                            home_email: record[i]["master_email"],
                            home_note: record[i]["master_note"],
                            home_bank_id: record[i]["master_bank_id"],
                            home_bank_norek: record[i]["master_bank_norek"],
                            home_bank_name: record[i]["master_bank_name"],
                            home_listadmin: record[i]["master_listadmin"],
                            home_listagen: record[i]["master_listagen"],
                            home_status: record[i]["master_status"],
                            home_status_css: record[i]["master_status_css"],
                            home_create: record[i]["master_create"],
                            home_update: record[i]["master_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_curr.length; i++) {
                listCurr = [
                    ...listCurr,
                    {
                        curr_id: record_curr[i]["curr_id"],
                    },
                ];
            }
            for (var i = 0; i < record_bank.length; i++) {
                listBank = [
                    ...listBank,
                    {
                        bank_category: record_bank[i]["catebank_name"],
                        bank_id: record_bank[i]["banktype_id"],
                    },
                ];
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
        listHome = [];
        listCurr = [];
        listBank = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {listCurr}
    {listBank}
    {totalrecord}/>
{/if}