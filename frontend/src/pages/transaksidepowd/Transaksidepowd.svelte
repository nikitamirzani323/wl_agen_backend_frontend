<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
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
                page: "PROVIDER-VIEW",
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
        const res = await fetch("/api/transdpwd", {
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
            let record_bank = json.listbank;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let amount_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(record[i]["transdpwd_tipedoc"] == "DEPOSIT"){
                        amount_css = "color:#3600ff;font-weight:bold;"
                    }else{
                        amount_css = "color:#ff0000;font-weight:bold;"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["transdpwd_id"],
                            home_tipedoc: record[i]["transdpwd_tipedoc"],
                            home_date: record[i]["transdpwd_date"],
                            home_info: record[i]["transdpwd_notebank"],
                            home_idmember: record[i]["transdpwd_idmember"],
                            home_nmmember: record[i]["transdpwd_nmmember"],
                            home_bank_in: record[i]["transdpwd_bank_in"],
                            home_bank_in_info: record[i]["transdpwd_bank_in_info"],
                            home_bank_out: record[i]["transdpwd_bank_out"],
                            home_bank_out_info: record[i]["transdpwd_bank_out_info"],
                            home_before: record[i]["transdpwd_before"],
                            home_after: record[i]["transdpwd_after"],
                            home_amount: record[i]["transdpwd_amount"],
                            home_amount_css: amount_css,
                            home_note: record[i]["transdpwd_note"],
                            home_status: record[i]["transdpwd_status"],
                            home_status_css: record[i]["transdpwd_status_css"],
                            home_create: record[i]["transdpwd_create"],
                            home_update: record[i]["transdpwd_update"],
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
        listHome = [];
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
    {totalrecord}/>
{/if}