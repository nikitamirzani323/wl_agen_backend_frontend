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
                page: "CATEBANK-VIEW",
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
        const res = await fetch("/api/catebank", {
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
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["catebank_id"],
                            home_name: record[i]["catebank_name"],
                            home_status: record[i]["catebank_status"],
                            home_status_css: record[i]["catebank_status_css"],
                            home_list: record[i]["catebank_list"],
                            home_create: record[i]["catebank_create"],
                            home_update: record[i]["catebank_update"],
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