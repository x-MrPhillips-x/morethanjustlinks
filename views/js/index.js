async function submitForm(e){
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const formJson = Object.fromEntries(formData.entries());
    
    const data = await getData('/login', {
      method: form.method,
      headers:{
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(formJson)
    });

    console.log("This is what we have");
    console.log({data});

    var error = "";
    var msg = "";
    var username = "";
    
    
    let entries = Object.entries(data);
    entries.map(([k,v] = entry) => {
      switch(k){
        case "error":
          error = v;
        case "msg":
          error = v;
        case "user":
          username = v;
      }
    });

    if (isError(data)) {
      console.log("there should not be an error")
      window.location.replace(`/`)
      return
    }

    if (isSubmittedLoginForm(data)) {
      window.location.replace(`/api/${username}/profile`);
      return
    }
    
}

async function getData(url,requestBody) {
  const resp = await fetch(url,requestBody);
  
  return resp.json();

}

function isError(data){
  let error = "";
  let entries = Object.entries(data);
  entries.map(([k,v] = entry) => {
    if(k=="error") {
      error = v
      console.log("there should not be an error 1")
      return
    }
  });

  return error != ""
}

function isSubmittedLoginForm(data) {
  let username = "";
  let entries = Object.entries(data);
  entries.map(([k,v] = entry) => {
    if(k=="user") {
      username = v
      console.log("there should not be an error 2")
      return
    }
  });

  return username != ""
}