<script>
  let home = false;
  let services = false;
  let project = false;

  const homeClicked = () => {
    home = true;
    services = false;
    project = false;
  };

  const servicesClicked = () => {
    home = false;
    services = true;
    project = false;
  };

  const projectClicked = () => {
    home = false;
    services = false;
    project = true;
  };
</script>

<nav>
  <h1 class="index">
    <a href="#home">JRC <span class="h1">Portfolio</span></a>
  </h1>
  <div class="navigation">
    <span>
      <a href="#home" class:home on:click={homeClicked}>Home</a>
    </span>
    <span>
      <a href="#Services" class:services on:click={servicesClicked}>Services</a>
    </span>
    <span>
      <a href="#Projects" class:project on:click={projectClicked}
        >Projects & Experiences</a
      >
    </span>
  </div>
  <div class="contact">
    <span class="first"><a href="#Contact">Hire me</a></span>
    <span class="second"><a href="#Contact">Get in touch</a></span>
  </div>
</nav>

<style>
  /* General styles for elements */
  .index {
    color: var(--heading1);
  }

  /* Apply styles when a link is active (clicked) */
  a:active {
    text-decoration: underline;
    color: var(--heading1);
  }

  /* Navigation bar styles */
  nav {
    position: fixed;
    top: 0;
    left: 0;
    padding: 10px 60px 0;
    display: flex;
    justify-content: space-between;
    font-family: Outfit, sans-serif;
    background: var(--nav-color1);
    height: 80px;
    width: 100vw;
    box-shadow: 0 2px 15px -6px black, 0 2px 0 -6px black;
  }

  /* Header 1 styles */
  h1 {
    font-size: var(--navh1);
    letter-spacing: 0.2rem;
    padding: 20px 10px 10px;
    cursor: pointer;
  }

  h1 a {
    text-decoration: none;
    color: var(--heading1);
  }

  /* Header 1 styles when class "h1" is applied */
  h1 .h1 {
    color: var(--heading2);
    letter-spacing: 0.1rem;
  }

  /* Navigation links styles */
  .navigation {
    display: flex;
    padding: 20px 10px 10px;
    justify-content: space-evenly;
    gap: 3rem;
  }

  .navigation span {
    align-items: center;
    justify-content: center;
    font-weight: bold;
    letter-spacing: 0.2rem;
    padding: 10px 10px 20px;
  }

  .navigation a {
    text-decoration: none;
    color: var(--heading1);
    font-size: var(--navul);
    text-align: center;
  }

  /* Contact section styles */
  .contact {
    display: flex;
    padding: 10px;
    justify-content: space-evenly;
  }

  .contact span {
    align-items: center;
    justify-content: center;
    font-weight: bold;
    letter-spacing: 0.1rem;
    padding: 10px 10px 20px;
  }

  .contact .second a {
    text-decoration: none;
    color: var(--heading2);
    font-size: var(--navul);
    text-align: center;
    margin-right: 20px;
    padding: 8px 20px;
    height: 40px;
    width: 150px;
    border: 2px solid var(--heading2);
    border-radius: 20px;
  }

  .contact .first a {
    background-color: var(--heading2);
    color: var(--heading1);
    font-size: var(--navul);
    text-align: center;
    margin-right: 20px;
    padding: 8px 20px;
    height: 40px;
    width: 150px;
    text-decoration: none;
    border-radius: 20px;
    border: 2px solid var(--heading2);
  }

  /* Add hover animation for links */
  .contact a:hover {
    animation: hover 3s;
  }
</style>
