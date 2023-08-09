<script>
    import {onMount} from 'svelte';
    import clickabled from './clickable.js';

    let token = localStorage.getItem("token");
    
    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "DASHBOARD-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp()
    let name = "word";
    let dom_element;
    let show = false;
    onMount(() => {
        console.log(dom_element)
        dom_element.addEventListener("click", () => {
            console.log("You clicked ", dom_element)
        })

        var toastTrigger = document.getElementById('liveToastBtn')
var toastLiveExample = document.getElementById('liveToast')
if (toastTrigger) {
  toastTrigger.addEventListener('click', function () {
    var toast = new bootstrap.Toast(toastLiveExample)

    toast.show()
  })
}
    })
</script>

<div class="container" style="margin-top: 100px;">
    <div class="row">
        <div class="col-lg-12">
            <h1 bind:this={dom_element}>Hai Semua</h1><br>

            <input bind:value={name}>
            <button type="button" on:click={() => show = !show}>Toggle</button>
            {#if show}
            <h2 use:clickabled={{name}}>Hai Semua h2</h2>
            {/if}
            <br>
            <br>

            <button 
                type="button" 
                class="btn btn-primary" id="liveToastBtn">Show live toast</button>
        </div>

    </div>
</div>  

<div class="position-fixed bottom-0 start-50 translate-middle-x" style="z-index: 11">
    <div id="liveToast" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
      <div class="toast-header">
        <img src="..." class="rounded me-2" alt="...">
        <strong class="me-auto">Bootstrap</strong>
        <small>11 mins ago</small>
        <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
      </div>
      <div class="toast-body">
        Hello, world! This is a toast message.
      </div>
    </div>
</div>