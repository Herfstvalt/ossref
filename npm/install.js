const { execSync } = require("child_process");
const fs = require("fs");
const path = require("path");
const https = require("https");

const REPO = "herfstvalt/ossref";
const BIN_DIR = path.join(__dirname, "bin");

function getPlatform() {
  const platform = process.platform;
  const arch = process.arch;

  const osMap = { darwin: "darwin", linux: "linux", win32: "windows" };
  const archMap = { x64: "amd64", arm64: "arm64" };

  const os = osMap[platform];
  const cpu = archMap[arch];

  if (!os || !cpu) {
    console.error(`Unsupported platform: ${platform}/${arch}`);
    process.exit(1);
  }

  return { os, cpu };
}

function getLatestVersion() {
  return new Promise((resolve, reject) => {
    https.get(
      `https://api.github.com/repos/${REPO}/releases/latest`,
      { headers: { "User-Agent": "ossref-npm" } },
      (res) => {
        if (res.statusCode === 302 || res.statusCode === 301) {
          const tag = res.headers.location.split("/").pop();
          resolve(tag);
          return;
        }
        let data = "";
        res.on("data", (chunk) => (data += chunk));
        res.on("end", () => {
          try {
            resolve(JSON.parse(data).tag_name);
          } catch {
            reject(new Error("Failed to get latest version"));
          }
        });
      }
    ).on("error", reject);
  });
}

async function install() {
  const { os, cpu } = getPlatform();

  let version;
  try {
    version = await getLatestVersion();
  } catch {
    console.error("Failed to fetch latest version. Skipping binary download.");
    process.exit(0);
  }

  const ext = os === "windows" ? "zip" : "tar.gz";
  const url = `https://github.com/${REPO}/releases/download/${version}/ossref_${os}_${cpu}.${ext}`;

  console.log(`Installing ossref ${version} (${os}/${cpu})...`);

  fs.mkdirSync(BIN_DIR, { recursive: true });

  const tmpFile = path.join(BIN_DIR, `ossref.${ext}`);

  try {
    execSync(`curl -sL "${url}" -o "${tmpFile}"`, { stdio: "inherit" });

    if (ext === "tar.gz") {
      execSync(`tar -xzf "${tmpFile}" -C "${BIN_DIR}"`, { stdio: "inherit" });
    } else {
      execSync(`unzip -o "${tmpFile}" -d "${BIN_DIR}"`, { stdio: "inherit" });
    }

    fs.unlinkSync(tmpFile);

    if (os !== "windows") {
      fs.chmodSync(path.join(BIN_DIR, "ossref"), 0o755);
    }

    console.log("✓ ossref installed successfully");
  } catch (err) {
    console.error("Failed to install ossref binary:", err.message);
    console.error("You can install manually from: https://github.com/herfstvalt/ossref/releases");
    process.exit(0);
  }
}

install();
