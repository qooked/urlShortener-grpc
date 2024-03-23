const form = document.getElementById("form");

form.onsubmit = async (e) => {
    e.preventDefault();
    const fd = new FormData(form);
    await fetch("/shorten", {
      method: "POST",
      body: fd,
    })
      .then(async (response) => {
        console.log(response.status);
        if (response.status === 200) {
          response.json().then((data) => {
            let shortenedUrl = data.shortened_url;
            alert(shortenedUrl);
          });
          return;
        }else{
          alert(await response.text())
          return;
        }
      })
      .catch((error) => {
        alert(error);
      });
  };
