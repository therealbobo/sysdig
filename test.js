import fs from 'node:fs';
import loadSysdig  from './build/userspace/sysdig/sysdig.js';
import loadCsysdig from './build/userspace/sysdig/csysdig.js';

const runSysdig = async (args) => {
	const Module = await loadSysdig();

	Module.FS.mkdir("./chisels");

	const chisel_dir = './userspace/sysdig/chisels/';

	fs.readdirSync(chisel_dir).map(filename => {
		const data = fs.readFileSync(chisel_dir + filename);
		Module.FS.writeFile('./chisels/' + filename, data);
	});

	args.map(arg => {
		if(arg.endsWith('.scap'))
		{
			const data = fs.readFileSync(arg);
			Module.FS.writeFile(arg, data);
		}
	});

	Module.callMain(args);

};

const runCsysdig = async (args) => {
	const Module = await loadCsysdig();

	Module.FS.mkdir("./chisels");

	const chisel_dir = './userspace/sysdig/chisels/';

	fs.readdirSync(chisel_dir).map(filename => {
		const data = fs.readFileSync(chisel_dir + filename);
		Module.FS.writeFile('./chisels/' + filename, data);
	});

	args.map(arg => {
		if(arg.endsWith('.scap'))
		{
			const data = fs.readFileSync(arg);
			Module.FS.writeFile(arg, data);
		}
	});

	Module.callMain(args);
};

var args = process.argv;
args.shift()
args.shift()
const program = args.shift();
if(program == 'sysdig')
{
	runSysdig(args);
}
else
{
	runCsysdig(args);
}
