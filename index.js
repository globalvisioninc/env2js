#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const { exec } = require("child_process");

let platform;
let arch;
let ext = "";

// TODO: support macOS and x32
if (process.platform.match(/win/)) {
  platform = "windows";
  ext = ".exe";
} else {
  platform = "linux";
}
if (process.arch.match(/x64/)) {
  arch = "amd64";
} else {
  arch = "x32";
}

let command = path.join(
  __dirname,
  "dist",
  `${platform}-${arch}`,
  `env2js${ext}`
);

if (!fs.existsSync(command)) {
  console.error(
    `Unsupported platform or architecture (${process.platform} / ${process.arch})`
  );
  process.exit(1);
}

const args = process.argv.slice(2);
if (args.length) {
  command += " " + args.join(" ");
}

exec(command, (err, stderr, stdout) => {
  if (err) {
    console.error(stderr.trimRight());
    return;
  }
  if (stdout.trim().length) {
    console.log(stdout.trimRight());
  }
});
