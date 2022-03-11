ubuntu20.04  disk extend 
Ubuntu 为主分区扩容 – 命令行

root@master:~# df -h  | grep  -v docker | grep -v pod
Filesystem      Size  Used Avail Use% Mounted on
udev            2.9G     0  2.9G   0% /dev
tmpfs           580M  9.7M  571M   2% /run
/dev/sda1       9.6G  8.1G  1.5G  85% /
tmpfs           2.9G     0  2.9G   0% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           2.9G     0  2.9G   0% /sys/fs/cgroup
/dev/loop0       62M   62M     0 100% /snap/core20/1328
/dev/loop1       62M   62M     0 100% /snap/core20/1361
/dev/sda15      105M  5.2M  100M   5% /boot/efi
/dev/loop2       68M   68M     0 100% /snap/lxd/21835
/dev/loop4       44M   44M     0 100% /snap/snapd/14978
/dev/loop3       68M   68M     0 100% /snap/lxd/22526
tmpfs           580M     0  580M   0% /run/user/1000
root@master:~#



root@master:~# lsblk  
NAME    MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
fd0       2:0    1    4K  0 disk
loop0     7:0    0 61.9M  1 loop /snap/core20/1328
loop1     7:1    0 61.9M  1 loop /snap/core20/1361
loop2     7:2    0 67.2M  1 loop /snap/lxd/21835
loop3     7:3    0 67.9M  1 loop /snap/lxd/22526
loop4     7:4    0 43.6M  1 loop /snap/snapd/14978
sda       8:0    0   20G  0 disk
├─sda1    8:1    0  9.9G  0 part /
├─sda14   8:14   0    4M  0 part
└─sda15   8:15   0  106M  0 part /boot/efi
sr0      11:0    1 1024M  0 rom  
root@master:~# parted /dev/sda
GNU Parted 3.3
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.
(parted) p                                                                
Model: VMware, VMware Virtual S (scsi)
Disk /dev/sda: 21.5GB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start   End     Size    File system  Name  Flags
14      1049kB  5243kB  4194kB                     bios_grub
15      5243kB  116MB   111MB   fat32              boot, esp
1      116MB   10.7GB  10.6GB  ext4

(parted) resizepart 1                                                     
Warning: Partition /dev/sda1 is being used. Are you sure you want to continue?
Yes/No? yes                                                               
End?  [10.7GB]? 21.5GB                                                    
(parted) q                                                                
Information: You may need to update /etc/fstab.

root@master:~# lsblk                                                      
NAME    MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
fd0       2:0    1    4K  0 disk
loop0     7:0    0 61.9M  1 loop /snap/core20/1328
loop1     7:1    0 61.9M  1 loop /snap/core20/1361
loop2     7:2    0 67.2M  1 loop /snap/lxd/21835
loop3     7:3    0 67.9M  1 loop /snap/lxd/22526
loop4     7:4    0 43.6M  1 loop /snap/snapd/14978
sda       8:0    0   20G  0 disk
├─sda1    8:1    0 19.9G  0 part /
├─sda14   8:14   0    4M  0 part
└─sda15   8:15   0  106M  0 part /boot/efi
sr0      11:0    1 1024M  0 rom  
root@master:~# resize2fs /dev/sda1
resize2fs 1.45.5 (07-Jan-2020)
Filesystem at /dev/sda1 is mounted on /; on-line resizing required
old_desc_blocks = 2, new_desc_blocks = 3
The filesystem on /dev/sda1 is now 5214459 (4k) blocks long.

root@master:~# df -h  | grep  -v docker | grep -v pod
Filesystem      Size  Used Avail Use% Mounted on
udev            2.9G     0  2.9G   0% /dev
tmpfs           580M  9.9M  571M   2% /run
/dev/sda1        20G  8.2G   12G  43% /
tmpfs           2.9G     0  2.9G   0% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           2.9G     0  2.9G   0% /sys/fs/cgroup
/dev/loop0       62M   62M     0 100% /snap/core20/1328
/dev/loop1       62M   62M     0 100% /snap/core20/1361
/dev/sda15      105M  5.2M  100M   5% /boot/efi
/dev/loop2       68M   68M     0 100% /snap/lxd/21835
/dev/loop4       44M   44M     0 100% /snap/snapd/14978
/dev/loop3       68M   68M     0 100% /snap/lxd/22526
tmpfs           580M     0  580M   0% /run/user/1000
root@master:~#



https://mikublog.com/tech/2458