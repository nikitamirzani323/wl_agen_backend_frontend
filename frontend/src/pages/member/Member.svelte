<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
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
        const res = await fetch("/api/member", {
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
                let credit_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(parseInt(record[i]["member_credit"])>0){
                        credit_css = "color:#3600ff;font-weight:bold;"
                    }else{
                        credit_css = "color:#ff0000;font-weight:bold;"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["member_id"],
                            home_ipaddress: record[i]["member_ipaddress"],
                            home_lastlogin: record[i]["member_lastlogin"],
                            home_timezone: record[i]["member_timezone"],
                            home_username: record[i]["member_username"],
                            home_name: record[i]["member_name"],
                            home_phone: record[i]["member_phone"],
                            home_email: record[i]["member_email"],
                            home_credit: record[i]["member_credit"],
                            home_credit_css: credit_css,
                            home_listbank: record[i]["member_listbank"],
                            home_status: record[i]["member_status"],
                            home_status_css: record[i]["member_status_css"],
                            home_create: record[i]["member_create"],
                            home_update: record[i]["member_update"],
                        },
                    ];
                }
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
    {listBank}
    {totalrecord}/>
{/if}