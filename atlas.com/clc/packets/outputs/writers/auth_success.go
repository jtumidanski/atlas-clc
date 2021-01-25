package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteAuthSuccess(accountId int, name string, gender byte, pic string) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.LoginStatus)
	w.WriteInt(0)
	w.WriteShort(0)
	w.WriteInt(uint32(accountId))
	w.WriteByte(gender)
	//boolean canFly = false;// Server.getInstance().canFly(client.getAccID());
	//writer.writeBool((YamlConfig.config.server.USE_ENFORCE_ADMIN_ACCOUNT || canFly) && client.getGMLevel() > 1);    // GM
	w.WriteBool(false)
	//writer.write(((YamlConfig.config.server.USE_ENFORCE_ADMIN_ACCOUNT || canFly) && client.getGMLevel() > 1) ? 0x80 : 0);  //
	// Admin Byte. 0x80,0x40,0x20.. Rubbish.
	w.WriteByte(0)
	// country code
	w.WriteByte(0)
	w.WriteAsciiString(name)
	w.WriteByte(0)
	// quiet ban
	w.WriteByte(0)
	// quiet ban timestamp
	w.WriteLong(0)
	// creation timestamp
	w.WriteLong(0)
	// 1 : remove the select the world you want to play in
	w.WriteInt(1)
	// 0 = Pin-System Enabled, 1 = Disabled
	w.WriteByte(1)
	// 0 = Register PIC, 1 = Ask for PIC, 2 = Disabled
	w.WriteByte(2)
	return w.Bytes()
}
